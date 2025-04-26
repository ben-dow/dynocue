import { Events, Window } from "@wailsio/runtime";
import { createContext, ReactNode, useContext, useEffect, useReducer } from 'react';
import { GetShowMetadata, GetSources } from "../../bindings/dynocue/cmd/dynocue/dynocueservice";
import { ShowMetadata, Sources } from "../../bindings/dynocue/pkg/model/models";


export interface Show {
    Metadata: ShowMetadata
    Sources: Sources
}

class defaultShow implements Show {
    Metadata: ShowMetadata;
    Sources: Sources;

    constructor() {
        this.Metadata = new ShowMetadata()
        this.Sources = {
            AudioSources: []
        }
    }
}


/**  Base Contexts for storing and updating show data */
const ShowContext = createContext<Show>(new defaultShow()); // Stores Show Information 
const ShowDispatchContext = createContext<React.ActionDispatch<[action: ShowUpdate]>>(() => { }); // Stores function for updating show


/** Type Definition for Children provided to ShowProvider */
type ContextProviderProps = {
    children?: ReactNode
}

/** ShowProvider must be placed at the root level of the component tree where a show can be used */
export function ShowProvider({ children }: ContextProviderProps) {
    const [show, dispatch] = useReducer(showReducer, new defaultShow())

    useEffect(() => {
        async function getDefaultShow() {
            const md = await GetShowMetadata()
            dispatch({ type: "METADATA", payload: md })

            const sources = await GetSources()
            dispatch({ type: "SOURCES", payload: sources })
        }
        Events.On("MODEL_UPDATE", (ev) => { dispatch(ev.data[0]) })
        getDefaultShow()
    }, [])

    useEffect(() => {
        Window.SetTitle(`${show.Metadata.Name === "" || show.Metadata.Name === undefined ? "Untitled" : show.Metadata.Name} - DynoCue`)
    }, [show.Metadata])

    return (
        <ShowContext.Provider value={show}>
            <ShowDispatchContext.Provider value={dispatch}>
                {children}
            </ShowDispatchContext.Provider>
        </ShowContext.Provider>
    )
}

/** Defines how a show can be updated */
interface ShowUpdate {
    type: string
    payload: any
}

/** Reducer Dispatch function for updating a show */
function showReducer(show: Show, action: ShowUpdate) {
    switch (action.type) {
        case "METADATA": {
            return {
                ...show,
                Metadata: action.payload,
            }
        }
        case "SOURCES": {
            return {
                ...show,
                Sources: action.payload
            }
        }
        default: {
            return show
        }
    }
}
//
// Show Hooks
//

/**
 * 
 * @returns The current show
 */
export function useShow(): Show {
    const show = useContext(ShowContext)
    return show
}

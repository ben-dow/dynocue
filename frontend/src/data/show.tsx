import { Events, Window } from "@wailsio/runtime";
import { createContext, ReactNode, useContext, useEffect, useReducer } from 'react';
import { GetShow } from "../../bindings/dynocue/cmd/dynocue/dynocueservice";
import { Show } from '../../bindings/dynocue/pkg/model/models';

/**  Base Contexts for storing and updating show data */
const ShowContext = createContext<Show>(new Show()); // Stores Show Information 
const ShowDispatchContext = createContext<React.ActionDispatch<[action: ShowUpdate]>>(() => { }); // Stores function for updating show


/** Type Definition for Children provided to ShowProvider */
type ContextProviderProps = {
    children?: ReactNode
}

/** ShowProvider must be placed at the root level of the component tree where a show can be used */
export function ShowProvider({ children }: ContextProviderProps) {
    const [show, dispatch] = useReducer(showReducer, new Show())

    useEffect(() => {
        async function getDefaultShow() {
            const show = await GetShow()
            if (show != null) {
                dispatch({ type: "SHOW", show: show })
            }
        }
        Events.On("MODEL_UPDATE", (ev) => { dispatch(ev.data[0]) })
        getDefaultShow()
    }, [])

    useEffect(() => {
        Window.SetTitle(`${show.name === "" || show.name === undefined ? "Untitled" : show.name} - DynoCue`)
    }, [show.name])

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
    show: Show | null
}

/** Reducer Dispatch function for updating a show */
function showReducer(show: Show, action: ShowUpdate) {
    switch (action.type) {
        case "SHOW": {
            if (action.show != null) {
                return action.show
            }
            return show
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
export function UseShow(): Show {
    const show = useContext(ShowContext)
    return show
}

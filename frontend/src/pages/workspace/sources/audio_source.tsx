import { Dialogs } from "@wailsio/runtime";
import { MRT_ColumnDef } from "mantine-react-table";
import { useMemo } from "react";
import { AddAudioSource } from "../../../../bindings/dynocue/cmd/dynocue/dynocueservice";
import { AudioSource } from "../../../../bindings/dynocue/pkg/model/models";
import { useShow } from "../../../data/show";
import { SourcesTable } from "./source_table";

export function AudioSources() {
    const show = useShow()

    const columns = useMemo<MRT_ColumnDef<AudioSource>[]>(() => [
        {
            accessorKey: "Label",
            header: "Label",
            enableEditing: true,
            mantineEditTextInputProps: ({ cell, row }) => ({
                onBlur: (event) => {
                }
            })
        },
    ], [])
    return (
        <div>
            <SourcesTable<AudioSource> columns={columns} data={show.Sources.AudioSources} addAction={AudioSourceAdd} addValue="Add Audio Source" deleteAction={() => { }} />
        </div>
    )

}

function AudioSourceAdd() {
    Dialogs.OpenFile({
        Title: "Open Audio Source",
        CanChooseFiles: true,
        AllowsMultipleSelection: false,
        Filters: [
            {
                DisplayName: "All Files",
                Pattern: "*"
            }
        ]
    }).then((path) => {
        if (path === "") {
            return
        }
        AddAudioSource(path)
    })


}
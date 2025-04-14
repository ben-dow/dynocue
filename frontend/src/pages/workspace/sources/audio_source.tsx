import { Dialogs } from "@wailsio/runtime";

export function AudioSources() {
    /*
      const show = UseShow()
  
      const columns = useMemo<MRT_ColumnDef<AudioSource>[]>(() => [
          {
              accessorKey: "Label",
              header: "Label",
              enableEditing: true,
              mantineEditTextInputProps: ({ cell, row }) => ({
                  onBlur: (event) => {
                      console.log(row)
                      UpdateAudioSourceLabel(row.original.Id, event.target.value)
                  }
              })
          },
      ], [])
  */
    /*
    return (
        <div>
            <SourcesTable<AudioSource> columns={columns} data={show.sourceList.AudioSources} addAction={AudioSourceAdd} addValue="Add Audio Source" deleteAction={DeleteAudioSource} />
        </div>
    )
        */

    return <div></div>
}

function AudioSourceAdd() {
    Dialogs.OpenFile({
        Title: "Open Audio Source",
        CanChooseFiles: true,
        AllowsMultipleSelection: true,
        Filters: [
            {
                DisplayName: "Audio Files",
                Pattern: "*.flac;*.mp3;*.m4a;*.wav;*.aac"
            },
            {
                DisplayName: "All Files",
                Pattern: "*"
            }
        ]
    })
}
import { ActionIcon, Box, Button, Flex, LoadingOverlay, Modal, NativeSelect, TextInput } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { IconPlayerPlay, IconTrash } from "@tabler/icons-react";
import { Dialogs } from "@wailsio/runtime";
import { MRT_ColumnDef } from "mantine-react-table";
import { useMemo, useState } from "react";
import { AddAudioSource, PlayAudioSource } from "../../../../bindings/dynocue/cmd/dynocue/dynocueservice";
import { AudioSource } from "../../../../bindings/dynocue/pkg/model/models";
import { useShow } from "../../../data/show";
import { SourcesTable } from "./source_table";

export function AudioSources() {
    const show = useShow()
    const [opened, { open, close }] = useDisclosure(false);

    const columns = useMemo<MRT_ColumnDef<AudioSource>[]>(() => [
        {
            header: "",
            id: "play",
            Cell: ({ cell }) => {
                return (
                    <Flex justify="center" gap={2}>
                        <ActionIcon color="green" onClick={() => { PlayAudioSource(cell.row.original.Id) }}>
                            <IconPlayerPlay />
                        </ActionIcon>
                    </Flex>
                )
            },
            maxSize: 10
        },
        {
            accessorKey: "Label",
            header: "Label",
            enableEditing: true,
            mantineEditTextInputProps: ({ cell, row }) => ({
                onBlur: (event) => {
                }
            })
        },
        {
            accessorKey: "Duration",
            header: "Duration",
            enableEditing: false,
            Cell: ({ cell }) => {
                const time: number = cell.getValue() as number
                const timeMs = Math.floor(time / 1000000)

                let seconds = Math.floor(timeMs / 1000)
                const hours = Math.floor(seconds / 3600)
                seconds = seconds % 3600
                const minutes = Math.floor(seconds / 60)
                seconds = seconds % 60

                const ms = String(timeMs % 1000).padStart(3, "0")
                const ss = String(seconds).padStart(2, "0")
                const mm = String(minutes).padStart(2, "0")
                const hh = String(hours).padStart(2, "0")
                const duration = `${hh}:${mm}:${ss}:${ms}`
                cell.getValue()
                return (<Box>
                    {hh}:{mm}:{ss}.{ms}
                </Box>)
            }
        },
        {
            header: "",
            id: "delete",
            Cell: ({ cell }) => {
                return (
                    <Flex justify="center" gap={2}>
                        <ActionIcon color="red" onClick={() => { }}>
                            <IconTrash />
                        </ActionIcon>
                    </Flex>
                )
            },
            maxSize: 20
        },
    ], [])

    console.log(show.Sources.AudioSources)
    return (
        <div>
            <AudioSourceAdd opened={opened} closer={close} />
            <SourcesTable<AudioSource> columns={columns} data={show.Sources.AudioSources} addAction={open} addValue="Add Audio Source" playAction={(id) => { PlayAudioSource(id) }} deleteAction={() => { }} />
        </div>
    )
}


interface AddAudioSourceProps {
    opened: boolean
    closer: () => void
}

function AudioSourceAdd(props: AddAudioSourceProps) {
    const [loaderVisible, loaderCtrl] = useDisclosure(false);

    const [label, setLabel] = useState("")
    const [codec, setCodec] = useState("mp3")
    const [file, setFile] = useState("")

    const reset = () => {
        setLabel("")
        setCodec("mp3")
        setFile("")
    }

    const opener = () => {
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
            setFile(path)
            if (label === "") {
                let fileWithExt = path.split('/').pop(); // Get the last part of the path (filename.ext)
                if (fileWithExt === undefined) {
                    fileWithExt = ""
                }
                const filename = fileWithExt.split('.').slice(0, -1).join('.'); // Remove the extension
                setLabel(filename)
            }
        })
    }

    const saver = () => {
        loaderCtrl.open()
        AddAudioSource(file, codec, label).then(() => {
            props.closer()
            reset()
        }).finally(() => {
            loaderCtrl.close()
        })
    }


    return (
        <Modal opened={props.opened} onClose={() => { props.closer(); reset() }} size={"lg"} title="Add New Audio Source" centered transitionProps={{ duration: 0 }}>
            <Box pos="relative">
                <LoadingOverlay visible={loaderVisible} zIndex={1000} overlayProps={{ radius: "sm", blur: 2 }} />
                <Flex gap={10} direction={"column"}>
                    <TextInput value={label} label={"Source Label"} onChange={(e) => { setLabel(e.target.value) }} />
                    <NativeSelect value={codec} label={"Storage Codec"} data={["mp3", "flac"]} onChange={(e) => { setCodec(e.target.value) }} />
                    <TextInput value={file} label="Select File" onClick={opener} readOnly={true} />
                    <Flex justify={"center"} gap={5}>
                        <Button color="red" onClick={() => { props.closer(); reset() }}>Cancel</Button>
                        <Button color="green" onClick={saver}>Add Source</Button>
                    </Flex>
                </Flex>
            </Box>
        </Modal >
    )

}
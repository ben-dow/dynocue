import {useNavigate} from "react-router";
import {
    ActionIcon,
    Box,
    Button,
    Card,
    Center,
    Checkbox,
    Divider,
    Flex,
    Group,
    Stack,
    TextInput,
    Title
} from "@mantine/core";
import {IconArrowBack, IconArrowLeft} from "@tabler/icons-react";

export default function ApplicationSettings() {
    const navigate = useNavigate()
    return (
        <Box className="h-screen bg-zinc-200">
            <Center className={"h-full"}>
                <Box className={"bg-white rounded p-2 w-2xl h-screen"}>
                    <Flex direction={"column"} gap={4}>
                        <Flex direction={"row"} >
                            <ActionIcon color={"black"}>
                                <IconArrowLeft onClick={()=>{navigate(-1)}}/>
                            </ActionIcon>
                            <Center className={"w-full"}>Application Settings</Center>
                        </Flex>
                        <Divider/>
                        <Box className={"h-full overflow-auto"}>
                            <Flex direction={"column"} gap={10}>
                                <TextInput label={"VLC"}  inputContainer={(children)=>
                                    <Group>
                                        <Box className={"w-4/5"}>
                                            {children}
                                        </Box>
                                        <Checkbox className={"w-1/6"} label={"Use PATH"}/>
                                    </Group>

                                }
                                />
                                <TextInput label={"FFMPEG"}  inputContainer={(children)=>
                                    <Group>
                                        <Box className={"w-4/5"}>
                                            {children}
                                        </Box>
                                        <Checkbox className={"w-1/6"} label={"Use PATH"}/>
                                    </Group>
                                }
                                />
                            </Flex>
                        </Box>
                    </Flex>
                </Box>
            </Center>
        </Box>
    )
}
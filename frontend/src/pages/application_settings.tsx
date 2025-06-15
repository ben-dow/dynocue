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
import { IconArrowLeft, IconDeviceFloppy, IconRestore, IconX} from "@tabler/icons-react";

export default function ApplicationSettings() {
    const navigate = useNavigate()
    return (
        <Box className="h-screen bg-zinc-200">
                <Stack justify="center" align="center">
                    <Card shadow="lg" padding="xl" radius={"md"} className={"w-2xl"} withBorder>
                    <Flex direction={"column"} gap={4}>
                        <Flex direction={"row"} >
                            <ActionIcon color={"black"}>
                                <IconArrowLeft onClick={()=>{navigate(-1)}}/>
                            </ActionIcon>
                            <Title className={"flex-1/3 text-center"}>Dependencies</Title>
                            <Flex direction={"row"} gap={4} className={"relative top-3"}>
                                <ActionIcon color={"green"}>
                                    <IconDeviceFloppy/>
                                </ActionIcon>
                                <ActionIcon color={"red"}>
                                    <IconX/>
                                </ActionIcon>
                                <ActionIcon color={"blue"}>
                                    <IconRestore/>
                                </ActionIcon>
                            </Flex>
                        </Flex>
                        <Divider/>
                        <Box className={"h-full overflow-auto"}>
                            <Flex direction={"column"} gap={10}>
                                <Title className={"w-full"}  order={4}>Application Dependencies</Title>
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

                                <TextInput label={"FFPLAY"}  inputContainer={(children)=>
                                    <Group>
                                        <Box className={"w-4/5"}>
                                            {children}
                                        </Box>
                                        <Checkbox className={"w-1/6"} label={"Use PATH"}/>
                                    </Group>
                                }
                                />

                                <TextInput label={"FFPROBE"}  inputContainer={(children)=>
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
                    </Card>
                </Stack>
        </Box>
    )
}
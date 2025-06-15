import { Button, Card, Center, Divider, Stack, Title } from "@mantine/core";
import { Dialogs } from "@wailsio/runtime";
import { useNavigate } from "react-router";
import {CreateNewAsHost, OpenAsHost} from "../../bindings/dynocue/cmd/dynocue/dynocueservice";


export default function Splash() {
    const navigate = useNavigate()

    const goToWorkspace = () => navigate("/workspace")


    return (
        <Center className="h-screen bg-zinc-200">
            <Stack justify="center" align="center">
                <Card shadow="lg" padding="xl" radius={"md"} withBorder>
                    <Stack gap={"xs"}>
                        <Title order={1}>DynoCue</Title>
                        <Divider />
                        <Stack gap={"sm"}>
                            <Button onClick={() => { CreateNewShow(goToWorkspace) }} size="lg" radius="md">New</Button>
                            <Button onClick={() => { OpenExistingShow(goToWorkspace) }} size="lg" radius="md">Open</Button>
                            <Button onClick={() => {  }} size="lg" radius="md">Connect</Button>
                            <Divider/>
                            <Button onClick={() => { navigate("/application_settings") }} size="lg" radius="md">Settings</Button>
                        </Stack>
                    </Stack>
                </Card>
            </Stack >
        </Center >
    )
}

function CreateNewShow(goToWorkspace: () => void) {
    Dialogs.SaveFile({
        CanCreateDirectories: true,
        CanChooseDirectories: true,
        CanChooseFiles: false,
        Title: "Create Dyno Cue Project",
    }).then((path) => {
        if (path === "") {
            return
        }
        CreateNewAsHost(path).then(() => { goToWorkspace() })
    })
}

function OpenExistingShow(goToWorkspace: () => void) {
    Dialogs.OpenFile({
        CanChooseDirectories: true,
        Title: "Open Dyno Cue Project",

    }).then((path) => {
        if (path === "") {
            return
        }
        OpenAsHost(path).then(() => { goToWorkspace() })
    })
}
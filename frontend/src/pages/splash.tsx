import { Button, Card, Center, Divider, Stack, Title } from "@mantine/core";
import { Dialogs } from "@wailsio/runtime";
import { useNavigate } from "react-router";
import { NewLocal, OpenLocal } from "../../bindings/dynocue/cmd/dynocue/dynocueservice";


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
                            <Button onClick={() => { NewShow(goToWorkspace) }} size="lg" radius="md">New</Button>
                            <Button onClick={() => { OpenShow(goToWorkspace) }} size="lg" radius="md">Open</Button>
                        </Stack>
                    </Stack>
                </Card>
            </Stack >
        </Center >
    )
}

function NewShow(goToWorkspace: () => void) {
    Dialogs.SaveFile({
        CanCreateDirectories: true,
        CanChooseDirectories: true,
        CanChooseFiles: false,
        Title: "Create Dyno Cue Project",
    }).then((path) => {
        if (path === "") {
            return
        }
        NewLocal(path).then(() => { goToWorkspace() })
    })
}

function OpenShow(goToWorkspace: () => void) {
    Dialogs.OpenFile({
        CanChooseDirectories: true,
        Title: "Open Dyno Cue Project",

    }).then((path) => {
        if (path === "") {
            return
        }
        OpenLocal(path).then(() => { goToWorkspace() })
    })
}
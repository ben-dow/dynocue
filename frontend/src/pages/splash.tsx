import { Button, Card, Center, Divider, Stack, Title } from "@mantine/core";
import { useNavigate } from "react-router";
import { NewLocal } from "../../bindings/dynocue/cmd/dynocue/dynocueservice";

export default function Splash() {
    const navigate = useNavigate()
    return (
        <Center className="h-screen bg-zinc-200">
            <Stack justify="center" align="center">
                <Card shadow="lg" padding="xl" radius={"md"} withBorder>
                    <Stack gap={"xs"}>
                        <Title order={1}>DynoCue</Title>
                        <Divider />
                        <Stack gap={"sm"}>
                            <Button onClick={() => { NewLocal().then(() => { navigate("/workspace") }) }} size="lg" radius="md">New</Button>
                            <Button size="lg" radius="md">Open</Button>
                            <Button size="lg" radius="md">Connect</Button>
                        </Stack>
                    </Stack>
                </Card>
            </Stack >
        </Center >
    )
}
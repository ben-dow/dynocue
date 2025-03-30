import { AppShell, Button, Stack } from "@mantine/core";
import { Outlet, useLocation, useNavigate } from "react-router";
import { ShowProvider, UseShow as useShow } from "../../data/show";

export default function Workspace() {
    return (
        <ShowProvider>
            <WorkspaceOutlet />
        </ShowProvider>
    )
}

function WorkspaceOutlet() {
    const show = useShow()

    return (
        <AppShell
            header={{ height: 60 }}
            navbar={{ width: 150, breakpoint: "xs" }}
            padding={"sm"}
        >
            <AppShell.Header className="">
            </AppShell.Header>
            <AppShell.Navbar>
                <Stack gap={"5"} p={2} className="bg-zinc-100 h-full">
                    <NavButton navPath="/workspace" label="Dashboard" />
                    <NavButton navPath="cues" label="Cues" />
                    <NavButton navPath="sources" label="Sources" />
                    <NavButton navPath="settings" label="Settings" />
                    <NavButton navPath="playback" label="Playback" />
                </Stack>
            </AppShell.Navbar>
            <AppShell.Main>
                <Outlet />
            </AppShell.Main>
        </AppShell>
    )
}


interface NavButtonProps {
    navPath: string
    label: string
}

function NavButton(props: NavButtonProps) {
    const navigate = useNavigate()
    const locate = useLocation()

    const active = locate.pathname.endsWith(props.navPath)

    return (
        <Button variant={active ? "filled" : "light"} radius="xs" onClick={() => { navigate(props.navPath) }} size="md" >{props.label}</Button>
    )
}
import { Tabs } from "@mantine/core";
import { AudioSources } from "./audio_source";

export default function Sources() {
    return (
        <Tabs defaultValue="audio">
            <Tabs.List>
                <Tabs.Tab value="audio">Audio</Tabs.Tab>
            </Tabs.List>
            <Tabs.Panel pt="sm" value="audio">
                <AudioSources />
            </Tabs.Panel>
        </Tabs>
    )
}



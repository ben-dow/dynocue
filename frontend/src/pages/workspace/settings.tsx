import { Stack, TextInput } from "@mantine/core";
import { SetShowName } from "../../../bindings/dynocue/cmd/dynocue/dynocueservice";
import { useShow } from "../../data/show";

export default function Settings() {
    const show = useShow()
    return (
        <Stack>
            <TextInput
                label="Show Name"
                value={show.Metadata.Name}
                onChange={(ev) => {
                    SetShowName(ev.currentTarget.value)
                }} />
        </Stack>
    )
}
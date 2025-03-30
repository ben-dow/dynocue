import { Stack, TextInput } from "@mantine/core";
import { SetShowName } from "../../../bindings/dynocue/cmd/dynocue/dynocueservice";
import { UseShow as useShow } from "../../data/show";

export default function Settings() {
    const show = useShow()
    return (
        <Stack>
            <TextInput label="Show Name" value={show.name || ""} onChange={(ev) => { SetShowName(ev.currentTarget.value) }} />
        </Stack>
    )
}
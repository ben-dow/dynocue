import { Input } from "@mantine/core";
import { SetShowName } from "../../../bindings/dynocue/cmd/dynocue/dynocueservice";
import { UseShow as useShow } from "../../data/show";

export default function Settings() {
    const show = useShow()

    return (
        <div>
            <Input value={show.name} onChange={(ev) => { SetShowName(ev.currentTarget.value) }} />
        </div>
    )
}
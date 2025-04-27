import { Button, Flex } from "@mantine/core";
import { MantineReactTable, MRT_ColumnDef, MRT_GlobalFilterTextInput, MRT_RowData } from "mantine-react-table";

export interface SourcesTableProps<T extends MRT_RowData> {
    columns: MRT_ColumnDef<T>[]
    data: T[]
    addAction: () => void
    addValue: string
    deleteAction: (id: string) => void
    playAction: (id: string) => void
}

export function SourcesTable<T extends MRT_RowData>(props: SourcesTableProps<T>) {
    return (
        <MantineReactTable
            columns={props.columns}
            data={props.data}
            enableColumnActions={false}
            enableColumnFilters={false}
            enableRowActions={false}
            enablePagination={false}
            enableSorting={false}
            enableEditing={true}
            editDisplayMode={"cell"}
            enableDensityToggle={false}
            enableFullScreenToggle={false}
            enableHiding={false}
            initialState={{ showColumnFilters: true, showGlobalFilter: true }}
            mantineTableProps={
                {
                    striped: 'odd',
                    withColumnBorders: true,
                    withRowBorders: true,
                    withTableBorder: true,
                }
            }
            mantineTableHeadCellProps={
                {
                    align: "center"
                }
            }
            mantineSearchTextInputProps={{
                placeholder: 'Search',
            }
            }
            displayColumnDefOptions={
                {
                    'mrt-row-actions': {
                        header: 'Actions', //change header text
                        size: 20, //make actions column wider
                    }
                }
            }
            renderTopToolbar={({ table }) => (
                <Flex p="15" justify={"space-between"}>
                    <Flex gap="xs">
                        <MRT_GlobalFilterTextInput table={table} />
                    </Flex>

                    <Flex>
                        <Button onClick={props.addAction}>{props.addValue}</Button>
                    </Flex>
                </Flex>
            )}
        />
    )
}
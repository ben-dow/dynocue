import { AppShell, Box } from '@mantine/core';
import { Route, Routes } from 'react-router';
import Cues from './pages/cues';

function App() {
    return (
        <Box>
            <AppShell
                header={{ height: 40 }}
                padding="md"
            >
                <AppShell.Header>
                    <div>Test</div>
                </AppShell.Header>

                <AppShell.Main>
                    <Routes>
                        <Route index element={<Cues />} />
                    </Routes>
                </AppShell.Main>
            </AppShell>
        </Box>
    )
}

export default App

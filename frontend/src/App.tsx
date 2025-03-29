import { Route, Routes } from 'react-router';
import Splash from './pages/splash';
import Settings from './pages/workspace/settings';
import Workspace from './pages/workspace/workspace';

function App() {
    return (
        <Routes>
            <Route path="/" element={<Splash />} />
            <Route path="/workspace" element={<Workspace />}>
                <Route index element={<div>Home</div>} />
                <Route path="cues" element={<div>Cues</div>} />
                <Route path="sources" element={<div>Sources</div>} />
                <Route path="settings" element={<Settings />} />
                <Route path="playback" element={<div>Playback</div>} />
            </Route>
        </Routes >
    )
}

export default App

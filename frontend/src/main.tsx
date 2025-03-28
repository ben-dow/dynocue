import { createTheme, MantineProvider } from '@mantine/core';
import '@mantine/core/styles.css';
import { createRoot } from 'react-dom/client';
import { BrowserRouter } from 'react-router';
import App from './App';
import './style.css';

const container = document.getElementById('root')

const root = createRoot(container!)

const theme = createTheme({
    /** Put your mantine theme override here */
});

root.render(
    <MantineProvider theme={theme}>
        <BrowserRouter>
            <App />
        </BrowserRouter>
    </MantineProvider>


)

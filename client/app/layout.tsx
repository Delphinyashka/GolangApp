import {MantineProvider} from '@mantine/core';
import {ReactNode} from "react";

export default function RootLayout({children}: { children: ReactNode }) {
    return (
        <html>
        <body>
        <MantineProvider withGlobalStyles withNormalizeCSS>
            {children}
        </MantineProvider>
        </body>
        </html>
    );
}
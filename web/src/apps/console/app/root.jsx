import { Links, LiveReload, Outlet, Scripts } from "@remix-run/react";
import { SSRProvider } from "react-aria"
import stylesUrl from "../../../index.css";
import consoleStyleUrl from "./styles/index.css";
import Container from "./pages/container";
export const links = () => [
    { rel: "stylesheet", href: stylesUrl },
    { rel: "stylesheet", href: consoleStyleUrl },
];

export default function App() {
    return (
        <html lang="en">
            <head>
                <meta charSet="utf-8" />
                <meta
                    name="viewport"
                    content="width=device-width,initial-scale=1"
                />
                <title>Remix: So great, it's funny!</title>
                <Links />
            </head>
            <body className="antialiased">
                <SSRProvider>
                    <Container>
                        <Outlet />
                    </Container>
                </SSRProvider>
                <Scripts />
            </body>
        </html>
    );
}
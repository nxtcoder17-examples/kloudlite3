import Container from "../pages/container"
import { Outlet, Link } from "@remix-run/react"
export default Console = ({ }) => {
    return <Container>
        <Outlet />
    </Container>
}
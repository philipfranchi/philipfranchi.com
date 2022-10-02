import "./Footer.css"
import { Icon } from "./Icon"

export default function Footer() {
    return <>
        <hr />
        <footer className="Footer">
            <Icon link="https://www.github.com" src="/github.png" />
            <Icon link="https://www.linkedin.com" src="/linkedin.png" />
            <Icon link="https://www.instagram.com" src="/instagram.png" />
        </footer>
    </>
}
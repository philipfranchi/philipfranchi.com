import "./Footer.css"
import { Icon } from "./Icon"

export default function Footer() {
    return <>
        <hr />
        <footer className="Footer">
            <Icon link="https://www.github.com/philipfranchi" src="/github.png" />
            <Icon link="https://www.linkedin.com/in/philip-franchi-pereira/" src="/linkedin.png" />
            <Icon link="https://www.instagram.com/philfranchi/" src="/instagram.png" />
        </footer>
    </>
}
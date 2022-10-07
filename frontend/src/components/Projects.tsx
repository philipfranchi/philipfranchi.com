import "./Projects.css";
import { ExternalLink } from "./ExternalLink";

export default function Projects() {
  return (
    <main className={"Projects"}>
      <h2>Projects</h2>

      <h3>Upcoming</h3>
      <hr></hr>
      <ul>
        <li>MOBI/EPUB converter</li>
        <li>6502 Emulator</li>
        <li>KMeans Color Pallete swapper in Go {"->"} WASM</li>
      </ul>

      <h3>Completed</h3>
      <hr></hr>
      <ul>
        <li>
          <ExternalLink link="https://www.github.com/philipfranchi/philipfranchi.net">
            This Website!
          </ExternalLink>
        </li>
      </ul>
    </main>
  );
}

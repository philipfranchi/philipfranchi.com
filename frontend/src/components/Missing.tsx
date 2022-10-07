import { Link } from "react-router-dom";

export default function Missing() {
  return (
    <main>
      <h2>404</h2>
      <p>
        Sorry, this page does not exist. Please check out my other content{" "}
        <Link to="/blog">here</Link>
      </p>
    </main>
  );
}

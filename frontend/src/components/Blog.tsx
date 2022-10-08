import Posts from "./Posts";

export default function Blog() {
  return (
    <>
      <h2>Blog</h2>
      <p>
        I've done some stuff and I've written about it here, in case you want to
        follow along. I leave some of the small details to the reader, but in
        general I like to be thorough. Sometimes I want to know how something
        works, and other times I'm more interested in getting what I want
        working and circling around to the unknowns later.
      </p>
      <Posts />
    </>
  );
}

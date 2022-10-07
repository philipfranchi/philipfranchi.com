import { useState, useEffect } from "react";
import { BlogPostMetadata } from "./api";
import { Link } from "react-router-dom";

const Post = () => {
  const [posts, setPosts] = useState<BlogPostMetadata[]>([]);

  useEffect(() => {
    const getPostMetadata = async () => {
      await fetch("/api/blog/")
        .then((response) => response.json())
        .then((data) => {
          console.log(data);
          setPosts(data);
        });
    };
    getPostMetadata();
  }, []);

  const metadataItems = posts
    .map((post) => post as BlogPostMetadata)
    .filter((post) => !!post.title && !!post.slug)
    .map((post) => (
      <li>
        <Link to={post.slug}>{post.title}</Link>
      </li>
    ));
  return <ul>{metadataItems}</ul>;
};
export default Post;

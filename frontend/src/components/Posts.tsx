import { useState, useEffect } from "react";
import { BlogPostMetadata } from "./api";
import { Link } from "react-router-dom";

import "./Posts.css";

const Post = () => {
  const [posts, setPosts] = useState<BlogPostMetadata[]>([]);

  useEffect(() => {
    const getPostMetadata = async () => {
      await fetch("/api/blog/")
        .then((response) => response.json())
        .then((data) => {
          setPosts(data);
        });
    };
    getPostMetadata();
  }, []);

  let metadataItems = posts
    .filter((post) => !!post.title && !!post.slug)
    .map(({ slug, title }) => (
      <Link key={slug} to={slug}>
        {title}
      </Link>
    ));
  metadataItems = [...metadataItems, ...metadataItems];
  return (
    <>
      <h2 className="PostsHeader">Posts</h2>
      <div className="Posts">{metadataItems}</div>
    </>
  );
};
export default Post;

import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import ReactMarkdown from "react-markdown";
import Missing from "./Missing";

const Post = () => {
  const postId = useParams().postId;
  const [markdown, setMarkdown] = useState("");
  const [shouldRedirect, setShouldRedirect] = useState(false);
  useEffect(() => {
    const getPost = async () => {
      await fetch("/api/blog/" + postId)
        .then((response) => {
          if (response.status === 404) {
            setShouldRedirect(true);
          }
          return response.json();
        })
        .then((data) => {
          setMarkdown(data.body);
          document.title = data.title;
        });
    };
    getPost();
  }, [postId]);
  if (shouldRedirect) {
    return <Missing />;
  }
  return <ReactMarkdown children={`${markdown}`} />;
};
export default Post;

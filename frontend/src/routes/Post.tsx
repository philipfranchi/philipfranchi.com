import { useParams } from "react-router-dom";

export default function Post() {
    let params = useParams();
    const postId = params.postId;
    console.log(postId);
    return (
        <h2>Post: {postId}</h2>
    );
  }
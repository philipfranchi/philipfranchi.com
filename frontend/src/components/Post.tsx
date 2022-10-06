import React, { useState, useEffect } from 'react';
import { useParams } from "react-router-dom";
import ReactMarkdown from 'react-markdown'

const Post = () => {
    const postId = useParams().postId;
    const [markdown, setMarkdown] = useState("");
    useEffect(() => {
        const getPost = async () => {
            await fetch(
                '/api/blog/' + postId,
            ).then((response) => response.text())
                .then((data) => setMarkdown(data));
        }
        getPost();
    }, [postId]);

    return <ReactMarkdown children={markdown} />

}
export default Post;
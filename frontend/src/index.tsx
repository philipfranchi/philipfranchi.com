import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';
import { BrowserRouter, Navigate, Routes, Route } from "react-router-dom";
import App from './App';
import Home from './routes/Home';
import Projects from './routes/Projects';
import Blog from './routes/Blog';
import Post from './routes/Post';
import Posts from './routes/Posts';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} > 
        <Route index element={<Home />} />
        <Route path="projects" element={<Projects />} /> 
          <Route path="blog" element={<Blog />}>
            <Route index element={<Posts />} />
            <Route path=":postId" element={<Post />} />
          </Route>
        </Route>
        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
      
    </BrowserRouter>
  </React.StrictMode>
);

// TODO
// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();

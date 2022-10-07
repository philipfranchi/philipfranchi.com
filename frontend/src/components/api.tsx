export interface BlogPostMetadata {
  title: string;
  slug: string;
  date: string;
  tags: string;
}

export interface BlogPost extends BlogPostMetadata {
  body: string;
}

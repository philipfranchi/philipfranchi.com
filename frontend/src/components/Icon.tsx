import "./Icon.css";
import { ExternalLink } from "./ExternalLink";

export interface IconProps {
  link: string;
  src: string;
  alt?: string;
}

export function Icon({ link, src, alt }: IconProps) {
  return (
    <ExternalLink link={link}>
      <img src={src} alt={alt ? alt : ""} />
    </ExternalLink>
  );
}

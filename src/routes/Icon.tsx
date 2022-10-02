import './Icon.css';

export interface IconProps {
    link: string,
    src: string,
    alt?: string
}

export function Icon({link, src, alt}: IconProps) {
    return <a className="Icon" href={link} target="_blank" rel="noreferrer">
        <img src={src}  alt={alt ? alt : ""}/>
    </a>
}
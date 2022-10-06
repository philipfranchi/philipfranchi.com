import {FC, ReactNode} from 'react';

export interface ExternalLinkProps {
    link: string,
    children: ReactNode,
}

export const ExternalLink: FC<ExternalLinkProps> = ({link, children}: ExternalLinkProps) => {
    return <a className="Icon" href={link} target="_blank" rel="noreferrer">{children}</a>
}

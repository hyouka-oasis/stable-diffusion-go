type Styles = Record<string, string>;

declare module '*.svg' {
  import React = require('react');

  export const ReactComponent: React.FC<React.SVGProps<SVGSVGElement>>;

  const content: string;
  export default content;
}

declare module '*.png' {
  const content: string;
  export default content;
}

declare module '*.jpg' {
  const content: string;
  export default content;
}

declare module '*.less' {
  const content: Styles;
  export default content;
}

declare module '*.css' {
  const content: Styles;
  export default content;
}

declare module '*.md';
declare module '*.svg';
declare module '*.json';
declare module '*.png';
declare module '*.jpg';
declare module '*.jpeg';
declare module '*.webp';
declare module '*.gif';
declare module '*.scss';
declare module '*.less';
declare module '*.ico';
declare module '*.json';
declare module '*.mp4';
declare module '*.mp3';
declare module '*.re3d';
declare module '*.glb';

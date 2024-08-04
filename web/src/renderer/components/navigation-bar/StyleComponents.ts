import styled from "styled-components";
import { navBarHeight } from "renderer/shared";

export const NavigationBarActionWrap = styled.div`
    height: 100%;
    display: flex;
    justify-content: flex-end;
    user-select: none;
    app-region: drag;
    outline: none;
    position: relative;

    .actions-container {
        display: flex;
        align-items: center;
        height: 100%;
        app-region: no-drag;

        .actions {
            display: flex;
            align-items: center;
            padding-left: 40px;
            height: 100%;

            > div {
                width: ${navBarHeight()}px;
                height: 100%;
                display: flex;
                align-items: center;
                justify-content: center;
                app-region: no-drag;

                &.baseAction:hover {
                    background: var(--main-bg5);
                }
            }

            > div:nth-child(1) {
                svg {
                    rect {
                        fill: currentColor;
                    }
                }
            }

            > div:nth-child(2) {
                svg {
                    rect {
                        stroke: currentColor;
                    }
                }
            }

            > div:nth-child(3) {
                svg {
                    path {
                        stroke: currentColor;
                    }
                }
            }
        }
    }
`;

export const NavigationBarWrap = styled.div<{ backgroundColor?: string; navigationTextColor?: string }>`
    width: 100%;
    height: ${navBarHeight()}px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    background-color: var(--main-bg1);
    app-region: drag;
    color: var(--color-text-secondary);
`;


export const NavigationBarContentWrap = styled.div`
    app-region: no-drag;
`;


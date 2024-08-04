// @ts-nocheck
import React, { cloneElement, createElement, forwardRef, isValidElement, useCallback, useEffect, useRef, } from "react";
import SmoothScrollbar from "smooth-scrollbar";
import type { Scrollbar } from "smooth-scrollbar/scrollbar";
import type { ScrollbarOptions, ScrollStatus } from "smooth-scrollbar/interfaces";
import type { OverscrollEffect, OverscrollOptions, } from "smooth-scrollbar/plugins/overscroll";

export interface ScrollbarPlugin extends Record<string, unknown> {
    overscroll?: Partial<Omit<OverscrollOptions, "effect">> & {
        effect?: OverscrollEffect;
    };
}

export type ScrollbarProps = Partial<ScrollbarOptions> &
    React.PropsWithChildren<{
        className?: string;
        style?: React.CSSProperties;
        plugins?: ScrollbarPlugin;
        onScroll?: (status: ScrollStatus, scrollbar: Scrollbar | null) => void;
        scrollbarClassName?: string;
    }>;

const ReactSmoothScrollbar = forwardRef<Scrollbar, ScrollbarProps>(
    function ReactSmoothScrollbar(
        { children, className, style, scrollbarClassName, ...restProps },
        ref
    ) {
        const mountedRef = useRef(false);
        const scrollbar = useRef<Scrollbar>(null!);
        const handleScroll = useCallback<(status: ScrollStatus) => void>(
            status => {
                if (typeof restProps.onScroll === "function") {
                    restProps.onScroll(status, scrollbar.current);
                }
            },
        [ restProps.onScroll ]
        );

        const containerRef = useCallback(node => {
            if (node instanceof HTMLElement) {
                (async () => {
                    if (restProps.plugins?.overscroll) {
                        const { default: OverscrollPlugin } = await import(
                            "smooth-scrollbar/plugins/overscroll"
                        );
                        SmoothScrollbar.use(OverscrollPlugin);
                    }
                    scrollbar.current = SmoothScrollbar.init(node, restProps);
                    scrollbar.current.addListener(handleScroll);
                })();
            }
        }, []);

        useEffect(() => {
            if (ref) {
                (ref as React.MutableRefObject<Scrollbar>).current = scrollbar.current;
            }
            return () => {
                if (scrollbar.current) {
                    scrollbar.current.removeListener(handleScroll);
                    scrollbar.current.destroy();
                }
            };
        }, []);

        useEffect(() => {
            if (mountedRef.current === true) {
                if (scrollbar.current) {
                    Object.keys(restProps).forEach(key => {
                        if (!(key in scrollbar.current.options)) {
                            return;
                        }

                        if (key === "plugins") {
                            Object.keys(restProps.plugins).forEach(pluginName => {
                                scrollbar.current.updatePluginOptions(
                                    pluginName,
                                    restProps.plugins[pluginName]
                                );
                            });
                        } else {
                            scrollbar.current.options[key] = restProps[key];
                        }
                    });

                    scrollbar.current.update();
                }
            } else {
                mountedRef.current = true;
            }
        }, [ restProps ]);

        if (isValidElement(children)) {
            return cloneElement(children as React.ReactElement, {
                ref: containerRef,
                className:
                    (children.props.className ? `${children.props.className} ` : "") +
                    className,
                style: {
                    ...style,
                    ...children.props.style,
                },
            });
        }

        return createElement(
            "div",
            {
                ref: containerRef,
                className: scrollbarClassName,
                style: {
                    ...style,
                    WebkitBoxFlex: 1,
                    msFlex: 1,
                    MozFlex: 1,
                    flex: 1,
                },
            },
            createElement(
                "div",
                {
                    className,
                },
                children
            )
        );
    }
);

export { ReactSmoothScrollbar };

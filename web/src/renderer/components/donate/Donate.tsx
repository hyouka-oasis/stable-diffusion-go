import styled, { keyframes } from "styled-components";
import qqPay from "renderer/assets/images/qq-pay.png";
import aliPay from "renderer/assets/images/ali-pay.png";
import wechatPay from "renderer/assets/images/wechat-pay.png";
import wechatDonate from "renderer/assets/images/wechat-donate.png";
import qqDonate from "renderer/assets/images/qq-donate.png";
import aliDonate from "renderer/assets/images/ali-donate.png";
import wechatInfo from "renderer/assets/images/wechat-info.png";
import DonateSvt from "renderer/assets/svg-com/donate.svg";
import CloseSvg from "renderer/assets/svg-com/nav-shut-down.svg";
import { useRef } from "react";

const showQR = keyframes`
    from {
        transform: rotateX(90deg);
    }
    8% {
        opacity: 1;
        transform: rotateX(-60deg);
    }
    18% {
        opacity: 1;
        transform: rotateX(40deg);
    }
    34% {
        opacity: 1;
        transform: rotateX(-28deg);
    }
    44% {
        opacity: 1;
        transform: rotateX(18deg);
    }
    58% {
        opacity: 1;
        transform: rotateX(-12deg);
    }
    72% {
        opacity: 1;
        transform: rotateX(9deg);
    }
    88% {
        opacity: 1;
        transform: rotateX(-5deg);
    }
    96% {
        opacity: 1;
        transform: rotateX(2deg);
    }
    to {
        opacity: 1;
    }
`;

const hideQR = keyframes`
    from {
    }
    20%,
    50% {
        transform: scale(1.08, 1.08);
        opacity: 1;
    }
    to {
        opacity: 0;
        transform: rotateZ(40deg) scale(0.6, 0.6);
    }
`;

const fadeIn = keyframes`
    from {
        opacity: 0;
    }
    to {
        opacity: 1;
    }
`;

const fadeOut = keyframes`
    from {
        opacity: 1;
    }
    to {
        opacity: 0;
    }
`;

const DonateWrap = styled.div`
    position: absolute;
    bottom: 10px;
    transition: all .3s;
    right: -290px;

    &:hover {
        right: 0;
    }

    #DonateText {
        pointer-events: none;
        position: absolute;
        font-size: 12px;
        width: 24px;
        height: 24px;
        color: #fff;
        background-size: 20px;
        border-radius: 35px;
        text-align: center;
        top: -15px;
        left: -10px;
        transform: rotatez(-15deg);
        z-index: 0;

        svg {
            width: 100%;
            height: 100%;
        }
    }

    #donateBox {
        background-color: #fff;
        border: 1px solid #ddd;
        border-radius: 6px;
        width: 300px;
        height: 28px;
        display: flex;
        align-items: center;
        justify-content: center;
        position: relative;
        z-index: 2;


        > span:first-child {
            border-start-start-radius: 6px;
            border-end-start-radius: 6px;
        }

        > span:last-child {
            border-start-end-radius: 6px;
            border-end-end-radius: 6px;
        }

        > span {
            width: 74px;
            text-align: center;
            border-left: 1px solid #ddd;
            background: rgba(204, 217, 220, 0.1) no-repeat center center;
            background-size: 45px;
            transition: all .3s;
            cursor: pointer;
            overflow: hidden;
            height: 28px;
            filter: grayscale(1);
            opacity: 0.5;
            display: flex;
            align-items: center;
            justify-content: center;

            &:hover {
                background-color: rgba(204, 217, 220, 0.3);
                filter: grayscale(0);
                opacity: 1;
            }

            img {
                width: 80%;
            }
        }
    }

    #QRBox {
        z-index: 1;
        display: none;
        perspective: 400px;
        position: fixed;
        background-color: rgba(255, 255, 255, 0.3);
        width: 100%;
        height: 100%;
        left: 0;
        top: 0;

        #MainBox {
            pointer-events: none;
            cursor: pointer;
            text-align: center;
            width: 200px;
            height: 300px;
            background: #fff no-repeat center center;
            background-size: 190px;
            border-radius: 6px;
            box-shadow: 0 2px 7px rgba(0, 0, 0, 0.3);
            opacity: 0;
            transition: all 1s ease-in-out;
            transform-style: preserve-3d;
            overflow: hidden;
            position: absolute;
            left: calc(50% - 100px);
            top: calc(50% - 150px);
        }

        #MainBox.showQR {
            animation-name: ${showQR};
            animation-duration: 3s;
            animation-timing-function: ease-in-out;
            animation-delay: 300ms;
            animation-iteration-count: 1;
            animation-fill-mode: forwards;
        }

        #MainBox.hideQR {
            opacity: 1;
            animation-name: ${hideQR};
            animation-duration: 0.5s;
            animation-timing-function: ease-in-out;
            animation-iteration-count: 1;
            animation-fill-mode: forwards;
        }

        .close-svg {
            cursor: pointer;
            position: absolute;
            left: calc(50% - -80px);
            top: calc(50% - 145px);

            > svg {
                fill: var(--main-bg5);

                path {
                    stroke: var(--main-bg5)
                }
            }
        }
    }

    #QRBox.fadeIn {
        display: block;
        animation: ${fadeIn} 300ms;
    }

    #QRBox.fadeOut {
        display: block;
        animation: ${fadeOut} 300ms;
    }
`;

const Donate = () => {
    const donateBoxRef = useRef<HTMLDivElement | null>(null);
    const qRBoxRef = useRef<HTMLDivElement | null>(null);
    const mainBox = useRef<HTMLDivElement | null>(null);

    const showDonateHandler = (id: string) => {
        let url = wechatInfo;
        if (id == "qq") {
            url = qqDonate;
        }
        if (id == "wechat") {
            url = wechatDonate;
        }
        if (id == "ali") {
            url = aliDonate;
        }
        if (!id) return;
        if (mainBox && mainBox.current) {
            mainBox.current.style.backgroundImage = "url(" + url + ")";
            if (qRBoxRef.current) {
                qRBoxRef.current.classList.add("fadeIn");
            }
            mainBox.current.classList.add("showQR");
        }
    };

    return (
        <DonateWrap>
            <div id="DonateText"><DonateSvt/></div>
            <div id="donateBox" className="list pos-f tr3" ref={donateBoxRef}>
                <span id="QQPay" onClick={() => showDonateHandler("qq")}>
                    <img src={qqPay} alt=""/>
                </span>
                <span id="AliPay" onClick={() => showDonateHandler("ali")}>
                    <img src={aliPay} alt=""/>
                </span>
                <span id="WeChat" onClick={() => showDonateHandler("wechat")}>
                    <img src={wechatPay} alt=""/>
                </span>
                <span id="me" onClick={() => showDonateHandler("me")}>
                    联系我
                </span>
            </div>
            <div id="QRBox" className="pos-f left-100" ref={qRBoxRef}>
                <div id="MainBox" ref={mainBox}></div>
                <div
                    className={"close-svg"} onClick={() => {
                        mainBox?.current?.classList.remove('showQR');
                        mainBox?.current?.classList.add('hideQR');
                        setTimeout(() => {
                            qRBoxRef?.current?.classList.remove('fadeIn');
                            mainBox?.current?.classList.remove('hideQR');
                        }, 600);
                    }}><CloseSvg/></div>
            </div>
        </DonateWrap>
    );
};

export default Donate;

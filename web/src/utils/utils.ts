//将base64转换为blob

const dataURLtoBlob = (dataurl: string): Blob => {
    const arr = dataurl.split(',');
    console.log(arr, "arr");
    const mime = arr?.[0]?.match(/:(.*?);/)?.[1];
    const bstr = atob(arr[1]);
    let n = bstr.length;
    const u8arr = new Uint8Array(n);
    while (n--) {
        u8arr[n] = bstr.charCodeAt(n);
    }
    return new Blob([ u8arr ], { type: mime });
};

//将blob转换为file
const blobToFile = (theBlob: Blob, fileName: string): File => {
    return new File([ theBlob ], fileName, { type: theBlob.type });
};

export {
    dataURLtoBlob,
    blobToFile
};

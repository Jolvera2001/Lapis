import { SetStateAction, useCallback, useState } from "react";
import ReactCodeMirror, { ViewUpdate } from "@uiw/react-codemirror";
import { markdown } from "@codemirror/lang-markdown";


function App() {
    const [value, setValue] = useState<string | undefined>("# Hello world!")
    const onChange = useCallback((val: SetStateAction<string | undefined>, ViewUpdate: any) => {
        console.log('val:', val );
        setValue(val);
    }, [])

    return <ReactCodeMirror value={value}  height="300px" extensions={[markdown()]} onChange={onChange} />;
}

export default App

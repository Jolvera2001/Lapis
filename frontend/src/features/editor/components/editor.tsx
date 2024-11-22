import { useCallback, useState } from "react";
import ReactCodeMirror from "@uiw/react-codemirror";
import { markdown } from "@codemirror/lang-markdown";
import { ViewUpdate } from "@uiw/react-codemirror";
import { Extension } from "@uiw/react-codemirror";

interface EditorProps {
    initialValue?: string;
    height?: string;
    width?: string;
    className?: string;
    onContentChange?: (content: string) => void;
}

function MarkdownEditor({
    initialValue = "# Hello World!",
    height = "",
    width = "",
    className = "",
    onContentChange,
}: EditorProps) {
    const [content, setContent] = useState<string>(initialValue);
    const extensions: Extension[] = [markdown()];

    const handleChange = useCallback(
        (value: string, ViewUpdate: ViewUpdate) => {
            setContent(value)
            onContentChange?.(value);

            // You can access additional information about the change
            // console.log('Selection:', viewUpdate.state.selection);
            // console.log('Doc length:', viewUpdate.state.doc.length);
        },
        [onContentChange]
    );

    return(
        <ReactCodeMirror
            value={content}
            height={height}
            width={width}
            className={className}
            extensions={extensions}
            onChange={handleChange}
            theme="light"
            style={{
                 height: '100%', 
                 width: '100%',
                }}
            indentWithTab={true}
            basicSetup={{
                lineNumbers: true,
                highlightActiveLineGutter: true,
                highlightActiveLine: true,
                foldGutter: true,
                dropCursor: true,
                allowMultipleSelections: true,
                indentOnInput: true,
                bracketMatching: true,
            }}
        />
    )
}

export default MarkdownEditor

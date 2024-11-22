import { Button, ButtonGroup, Container, IconButton, Stack, Typography } from "@mui/joy"
import MarkdownEditor from "../components/editor"
import { SimpleTreeView, TreeItem } from "@mui/x-tree-view"
import { useState } from "react"
import { Add } from "@mui/icons-material"

function EditorPage() {
    const [currentFolder, setCurrentFolder] = useState("test")

    return(
        <Stack 
            direction="row"
            spacing={0}
            sx={{
                position: 'fixed',
                height: '100vh',
                width: '100vw',
                overflow: 'hidden',
              }}
        >
            <Stack
                spacing={1}
                sx={{
                    width: '240px',
                    height: '100%',
                    borderRight: '1px solid',
                    borderColor: 'divider',
                    overflow: 'auto',
                    padding: '5px'
                }}
            >
                <Stack
                    direction="row"
                    spacing={2}
                >
                    <Typography
                        level="h3"
                    >
                        {currentFolder}
                    </Typography> 
                    <Button startDecorator={<Add />}>New Folder</Button>
                </Stack>
                <SimpleTreeView>
                    <TreeItem label="someLabel" itemId="1" />
                </SimpleTreeView>
            </Stack>
            <Stack
                spacing={0}
                sx={{
                    flex: 1,
                    height: '100%',
                    overflow: 'hidden',
                }}
            >
                <MarkdownEditor
                    height="100%"
                    width="100%" 
                />
            </Stack>
        </Stack>
    )
}

export default EditorPage
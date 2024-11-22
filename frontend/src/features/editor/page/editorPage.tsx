import { Container, Stack } from "@mui/joy"
import MarkdownEditor from "../components/editor"
import { SimpleTreeView, TreeItem } from "@mui/x-tree-view"

function EditorPage() {
    return(
        <Stack 
            direction="row"
            spacing={0}
            sx={{
                position: 'fixed', // Fix the position
                height: '100vh',
                width: '100vw',
                overflow: 'hidden',
              }}
        >
            <Stack
                sx={{
                    width: '240px', // Fixed width for sidebar
                    height: '100%',
                    borderRight: '1px solid',
                    borderColor: 'divider',
                    overflow: 'auto', // Allow scrolling in tree view
                }}
            >
                <SimpleTreeView>
                    <TreeItem label="someLabel" itemId="1" />
                </SimpleTreeView>
            </Stack>
            <Stack
                sx={{
                    flex: 1, // Take remaining space
                    height: '100%',
                    overflow: 'hidden', // Important for editor
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
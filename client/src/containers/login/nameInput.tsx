import React from 'react'
import TextField from '@material-ui/core/TextField';

export default function NameInput(props: {id: string; name: string; label: string; autoComplete: string}) {
    return(
        <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id={props.id}
            label={props.label}
            name={props.name}
            autoComplete={props.autoComplete}
            autoFocus
        />
    )
}
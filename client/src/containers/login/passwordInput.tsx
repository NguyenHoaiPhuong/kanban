import React from 'react'
import TextField from '@material-ui/core/TextField';

export default function PasswordInput(props: { name: string; label: string; type: string; id: string; autoComplete: boolean; }) {
    let pass: string = "";
    if (props.autoComplete) {
        pass = "current-password";
    }
    return(        
        <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name={props.name}
            label={props.label}
            type={props.type}
            id={props.id}
            autoComplete={pass}
        />        
    )
}
import React from 'react'
import Button from '@material-ui/core/Button';
import useStyles from './submitButtonStyle'

export default function SubmitButton(props: {content: string}) {
    const classes = useStyles();
    return (
        <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
        >
            {props.content}
        </Button>
    )
}
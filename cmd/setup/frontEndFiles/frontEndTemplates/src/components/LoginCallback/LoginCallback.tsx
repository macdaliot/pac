import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import { Redirect } from 'react-router-dom';
import { connect } from 'react-redux';

interface Props {
    setToken(token: string): void;
    location: any;
}
export class LoginCallbackComponent extends React.Component<Props> {
    constructor(props: any) {
        super(props);
    }
    componentDidMount() {
        let token = this.props.location.search.slice(1);
        this.props.setToken(token);
    }

    render() {
        return <Redirect to="/" />;
    }
}

const mapState = state => ({});
const mapDispatch = dispatch => ({
    setToken: token => dispatch({ type: 'JWT_RECEIVED', token: token }) 
});
export default connect(mapState, mapDispatch)(hot(LoginCallbackComponent))
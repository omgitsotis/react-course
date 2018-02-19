import React, {Component} from 'react';
import ChannelSection from './channels/ChannelSection.jsx';
import Socekt from './socket.js';

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            channels: [],
            activeChannel: {},
            connected: false
        }
    }

    componentDidMount() {
        let socket = this.socket = new Socket();
        socket.on('connect', this.onConnect.bind(this));
        socket.on('disconnect', this.onDisconnect.bind(this));
        socket.on('channel add', this.onAddChannel.bind(this));

    }

    onConnect() {
        this.setState({connected: true});
        this.socket.emit('channel subscribe');
    }

    onDisconnect() {
        this.setState({connected: false});
    }

    onAddChannel(channel) {
        let {channels} = this.state;
        channels.push(channel);
        this.setState({channels});
    }

    addChannel(name) {
        this.socket.emit('channel add', {name});
    }

    setChannel(activeChannel) {
        this.setState({activeChannel});
    }

    render() {
        return (
            <div className='app'>
                <div className="nav">
                    <ChannelSection
                        {...this.state}
                        addChannel={this.addChannel.bind(this)}
                        setChannel={this.setChannel.bind(this)}
                        />
                </div>
            </div>
        )
    }
}

export default App

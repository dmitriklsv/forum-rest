import './style.css'
import { Header } from './header'

class App {
    init() {
        this.render()
    }

    render() {
        const root = document.getElementById("root")
        
        const header = new Header()
        const headerMarkup = header.getMarkup()
        root.appendChild(headerMarkup)
        
    }
}

let app = new App()
app.init()
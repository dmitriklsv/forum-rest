class Header {
    fragment: DocumentFragment

    constructor(HTMLFragment: DocumentFragment) {
        this.fragment = HTMLFragment
    }

    getMarkup() {
        const markup: any = `
            <div class="header">
                <h2 class="header-title">forum</h2>
                <a class="header-profile" href="javascript:;" onclick="onClickProfile()">Profile</a>
            </div>
        `

        return this.fragment.appendChild(markup)
    }
} 

export default Header
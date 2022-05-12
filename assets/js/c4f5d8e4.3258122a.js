"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[195],{2280:function(e,n,t){t.r(n),t.d(n,{default:function(){return v}});var l=t(7294),a=t(6010),r=t(2600),c=t(2263),i="heroBanner_qdFl",o="codeblock_gCdh",m="example_HoTd",s=t(7462),u=t(9960),d="features_t9lD",g="buttons_wjNP",h=[{title:"Run Actions Before Sending Signal",description:l.createElement(l.Fragment,null,"Run a special action before the debouncer sends signal with ",l.createElement("code",null,"debouncer.Do(specialFunc)"),". The debouncer first invokes ",l.createElement("code",null,"specialFunc"),", then sends signal to invoke triggered function after ",l.createElement("code",null,"wait")," time.")},{title:"Control Debouncer Lifecycle",description:l.createElement(l.Fragment,null,"Cancel debouncer from invoking the triggered function at any time with ",l.createElement("code",null,"debouncer.Cancel()"),". Send a signal to the debouncer again when you want to restart it.")},{title:"Update Debouncer After Sending Signal",description:l.createElement(l.Fragment,null,"Debouncer allows replacing the triggered function and the timer after the signal was sent. The new timer take effect in the next ",l.createElement("code",null,"SendSignal()"),".")},{title:"Notify the Caller When Triggered Func Finishes",description:l.createElement(l.Fragment,null,"Debouncer allows sending a signal via the ",l.createElement("code",null,"Done()")," channel to the caller to let it knows the triggered func has been executed successfully.")}];function f(e){var n=e.title,t=e.description;return l.createElement("div",{className:(0,a.Z)("col col--3")},l.createElement("div",{className:"text--left padding-horiz--lg"},l.createElement("h3",null,n),l.createElement("p",null,t)))}function b(){return l.createElement("section",{className:d},l.createElement("div",{className:"container"},l.createElement("div",{className:"row"},h.map((function(e,n){return l.createElement(f,(0,s.Z)({key:n},e))}))),l.createElement("div",{className:"row"},l.createElement("div",{className:(0,a.Z)("col col--12")},l.createElement("div",{className:g},l.createElement(u.Z,{className:"button button--primary button--lg",to:"https://github.com/vnteamopen/godebouncer"},"Documentation"))))))}var E=t(8066);function p(){var e=(0,c.Z)().siteConfig;return l.createElement("header",{className:(0,a.Z)("hero hero--primary",i)},l.createElement("div",{className:"container"},l.createElement("h1",{className:"hero__title"},e.title),l.createElement("p",{className:"hero__subtitle"},e.tagline),l.createElement(l.Fragment,null,l.createElement(E.Z,{className:o,language:"bash"},"go get -u github.com/vnteamopen/godebouncer"))))}function v(){(0,c.Z)().siteConfig;return l.createElement(r.Z,{title:"GoDebouncer"},l.createElement(p,null),l.createElement("main",null,l.createElement("div",{className:m},l.createElement(E.Z,{language:"go"},'package main\n\nimport (\n\t"fmt"\n\t"time"\n\n\t"github.com/vnteamopen/godebouncer"\n)\n\nfunc main() {\n\tdebouncer := godebouncer.New(5 * time.Second).WithTriggered(func() {\n\t\tfmt.Println("Trigger") // Triggered func will be called after 5 seconds from last SendSignal().\n\t})\n\n\tfmt.Println("Action 1")\n\tdebouncer.SendSignal()\n\n\ttime.Sleep(1 * time.Second)\n\n\tfmt.Println("Action 2")\n\tdebouncer.SendSignal()\n\n\t// After 5 seconds, the trigger will be called.\n\t// Previous `SendSignal()` will be ignored to trigger the triggered function.\n\t<-debouncer.Done()\n}\n')),l.createElement(b,null)))}}}]);
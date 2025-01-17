package quickstart

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	sdkStyle             = lipgloss.NewStyle().PaddingLeft(4)
	selectedSdkItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
)

const (
	clientSideSDK = "client"
	mobileSDK     = "mobile"
	serverSideSDK = "server"
)

type chooseSDKModel struct {
	help          help.Model
	helpKeys      keyMap
	list          list.Model
	selectedIndex int
	selectedSDK   sdkDetail
}

func NewChooseSDKModel(selectedIndex int) tea.Model {
	l := list.New(sdksToItems(), sdkDelegate{}, 30, 9)
	l.Title = "Select your SDK:\n"
	// reset title styles
	l.Styles.Title = lipgloss.NewStyle()
	l.Styles.TitleBar = lipgloss.NewStyle()
	l.SetShowHelp(false)
	l.SetShowPagination(true)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false) // TODO: try to get filtering working

	return chooseSDKModel{
		help: help.New(),
		helpKeys: keyMap{
			Back:          BindingBack,
			CursorUp:      BindingCursorUp,
			CursorDown:    BindingCursorDown,
			PrevPage:      BindingPrevPage,
			NextPage:      BindingNextPage,
			GoToStart:     BindingGoToStart,
			GoToEnd:       BindingGoToEnd,
			ShowFullHelp:  BindingShowFullHelp,
			CloseFullHelp: BindingCloseFullHelp,
			Quit:          BindingQuit,
		},
		list:          l,
		selectedIndex: selectedIndex,
	}
}

// Init sends commands when the model is created that will:
// * select an SDK if it's already been selected
func (m chooseSDKModel) Init() tea.Cmd {
	return selectedSDK(m.selectedIndex)
}

func (m chooseSDKModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, pressableKeys.Enter):
			i, ok := m.list.SelectedItem().(sdkDetail)
			if ok {
				m.selectedSDK = i
				m.selectedSDK.index = m.list.Index()
				cmd = chooseSDK(m.selectedSDK)
			}
		case key.Matches(msg, m.helpKeys.CloseFullHelp):
			m.help.ShowAll = !m.help.ShowAll
		default:
			m.list, cmd = m.list.Update(msg)
		}
	case selectedSDKMsg:
		m.list.Select(msg.index)
	}

	return m, cmd
}

func (m chooseSDKModel) View() string {
	return m.list.View() + footerView(m.help.View(m.helpKeys), nil)
}

type sdkDetail struct {
	canonicalName string
	displayName   string
	index         int
	kind          string
	url           string // custom URL if it differs from the other SDKs
}

func (s sdkDetail) FilterValue() string { return "" }

var SDKs = []sdkDetail{
	{
		canonicalName: "react",
		displayName:   "React",
		kind:          clientSideSDK,
		url:           "https://github.com/launchdarkly/react-client-sdk/tree/main/examples/typescript",
	},
	{canonicalName: "node-server", displayName: "Node.js (server-side)", kind: serverSideSDK},
	{canonicalName: "python", displayName: "Python", kind: serverSideSDK},
	{canonicalName: "java", displayName: "Java", kind: serverSideSDK},
	{canonicalName: "dotnet-server", displayName: ".NET (server-side)", kind: serverSideSDK},
	{canonicalName: "js", displayName: "JavaScript", kind: clientSideSDK},
	{
		canonicalName: "vue",
		displayName:   "Vue",
		kind:          clientSideSDK,
		url:           "https://github.com/launchdarkly/vue-client-sdk/tree/main/example",
	},
	{canonicalName: "ios-swift", displayName: "iOS", kind: mobileSDK},
	{canonicalName: "go", displayName: "Go", kind: serverSideSDK},
	{canonicalName: "android", displayName: "Android", kind: mobileSDK},
	{
		canonicalName: "react-native",
		displayName:   "React Native",
		kind:          mobileSDK,
		url:           "https://github.com/launchdarkly/js-core/tree/main/packages/sdk/react-native/example",
	},
	{canonicalName: "ruby", displayName: "Ruby", kind: serverSideSDK},
	{canonicalName: "flutter", displayName: "Flutter", kind: mobileSDK},
	{canonicalName: "dotnet-client", displayName: ".NET (client-side)", kind: clientSideSDK},
	{canonicalName: "erlang", displayName: "Erlang", kind: serverSideSDK},
	{canonicalName: "rust", displayName: "Rust", kind: serverSideSDK},
	{
		canonicalName: "c-client",
		displayName:   "C/C++ (client-side)",
		kind:          clientSideSDK,
		url:           "https://github.com/launchdarkly/cpp-sdks/tree/main/examples/hello-cpp-client",
	},
	{canonicalName: "roku", displayName: "Roku", kind: clientSideSDK},
	{canonicalName: "node-client", displayName: "Node.js (client-side)", kind: clientSideSDK},
	{
		canonicalName: "c-server",
		displayName:   "C/C++ (server-side)",
		kind:          serverSideSDK,
		url:           "https://github.com/launchdarkly/cpp-sdks/tree/main/examples/hello-cpp-server",
	},
	{
		canonicalName: "lua-server",
		displayName:   "Lua",
		kind:          serverSideSDK,
		url:           "https://github.com/launchdarkly/lua-server-sdk/tree/main/examples/hello-lua-server",
	},
	{canonicalName: "haskell-server", displayName: "Haskell", kind: serverSideSDK},
	{canonicalName: "php", displayName: "PHP", kind: serverSideSDK},
}

func sdksToItems() []list.Item {
	items := make([]list.Item, len(SDKs))
	for i, sdk := range SDKs {
		items[i] = list.Item(sdk)
	}

	return items
}

type sdkDelegate struct{}

func (d sdkDelegate) Height() int                             { return 1 }
func (d sdkDelegate) Spacing() int                            { return 0 }
func (d sdkDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d sdkDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(sdkDetail)
	if !ok {
		return
	}

	fn := sdkStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedSdkItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(fmt.Sprintf("%d. %s", index+1, i.displayName)))
}

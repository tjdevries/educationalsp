package educationlsp_test

import (
	"educationlsp"
	"testing"
)

func TestDecodeInitializeMessage(t *testing.T) {
	msg := `{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "initialize",
  "params": {
    "processId": 12345,
    "rootPath": "/path/to/workspace",
    "rootUri": "file:///path/to/workspace",
    "capabilities": {
      "workspace": {
        "didChangeConfiguration": {
          "dynamicRegistration": true
        }
      },
      "textDocument": {
        "synchronization": {
          "dynamicRegistration": true,
          "willSave": true,
          "willSaveWaitUntil": true,
          "didSave": true
        }
      }
    }
  }
}
`
	encoded := educationlsp.EncodeMessage(msg)

	decoded, err := educationlsp.DecodeMessage(encoded)
	if err != nil {
		t.Fatal(err)
	}

	if decoded.(educationlsp.InitializeMessage).Method != "initialize" {
		t.Fatalf("\n%+v\n", decoded)
	}
}

func TestRealDecode(t *testing.T) {
	msg := "Content-Length: 3487\r\n\r\n" + `{"id":1,"method":"initialize","params":{"trace":"off","rootPath":null,"workspaceFolders":null,"processId":3073154,"rootUri":null,"clientInfo":{"name":"Neovim","version":"0.10.0-dev+gb76a01055"},"capabilities":{"workspace":{"configuration":true,"semanticTokens":{"refreshSupport":true},"applyEdit":true,"workspaceFolders":true,"didChangeWatchedFiles":{"dynamicRegistration":true,"relativePatternSupport":true},"symbol":{"dynamicRegistration":false,"symbolKind":{"valueSet":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26]}},"workspaceEdit":{"resourceOperations":["rename","create","delete"]},"didChangeConfiguration":{"dynamicRegistration":false},"inlayHint":{"refreshSupport":true}},"window":{"showMessage":{"messageActionItem":{"additionalPropertiesSupport":false}},"showDocument":{"support":true},"workDoneProgress":true},"general":{"positionEncodings":["utf-16"]},"textDocument":{"semanticTokens":{"serverCancelSupport":false,"multilineTokenSupport":false,"overlappingTokenSupport":true,"tokenTypes":["namespace","type","class","enum","interface","struct","typeParameter","parameter","variable","property","enumMember","event","function","method","macro","keyword","modifier","comment","string","number","regexp","operator","decorator"],"dynamicRegistration":false,"requests":{"full":{"delta":true},"range":false},"tokenModifiers":["declaration","definition","readonly","static","deprecated","abstract","async","modification","documentation","defaultLibrary"],"formats":["relative"],"augmentsSyntaxTokens":true},"rangeFormatting":{"dynamicRegistration":true},"codeAction":{"dynamicRegistration":true,"dataSupport":true,"codeActionLiteralSupport":{"codeActionKind":{"valueSet":["","quickfix","refactor","refactor.extract","refactor.inline","refactor.rewrite","source","source.organizeImports"]}},"resolveSupport":{"properties":["edit"]},"isPreferredSupport":true},"callHierarchy":{"dynamicRegistration":false},"synchronization":{"dynamicRegistration":false,"didSave":true,"willSave":true,"willSaveWaitUntil":true},"documentSymbol":{"symbolKind":{"valueSet":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26]},"dynamicRegistration":false,"hierarchicalDocumentSymbolSupport":true},"documentHighlight":{"dynamicRegistration":false},"signatureHelp":{"dynamicRegistration":false,"signatureInformation":{"activeParameterSupport":true,"documentationFormat":["markdown","plaintext"],"parameterInformation":{"labelOffsetSupport":true}}},"formatting":{"dynamicRegistration":true},"rename":{"dynamicRegistration":true,"prepareSupport":true},"diagnostic":{"dynamicRegistration":false},"typeDefinition":{"linkSupport":true},"publishDiagnostics":{"tagSupport":{"valueSet":[1,2]},"relatedInformation":true,"dataSupport":true},"implementation":{"linkSupport":true},"declaration":{"linkSupport":true},"references":{"dynamicRegistration":false},"hover":{"dynamicRegistration":true,"contentFormat":["markdown","plaintext"]},"definition":{"linkSupport":true,"dynamicRegistration":true},"completion":{"completionItem":{"snippetSupport":false,"documentationFormat":["markdown","plaintext"],"deprecatedSupport":false,"preselectSupport":false,"commitCharactersSupport":false},"contextSupport":false,"dynamicRegistration":false,"completionItemKind":{"valueSet":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25]}},"inlayHint":{"dynamicRegistration":true,"resolveSupport":{"properties":["textEdits","tooltip","location","command"]}}}}},"jsonrpc":"2.0"}`

	decoded, err := educationlsp.DecodeMessage(msg)
	if err != nil {
		t.Fatal(err)
	}

	if decoded.(educationlsp.InitializeMessage).Method != "initialize" {
		t.Fatalf("\n%+v\n", decoded)
	}
}

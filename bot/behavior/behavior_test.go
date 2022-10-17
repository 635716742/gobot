package behavior

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var compose = `
<behavior>
  <id>14bcccc7-f3e0-41db-b4e1-df8ac960f178</id>
  <ty>RootNode</ty>
  <pos>
    <x>0</x>
    <y>0</y>
  </pos>
  <children>
    <id>63014028-4013-4474-b0e2-956812940859</id>
    <ty>LoopNode</ty>
    <pos>
      <x>-5</x>
      <y>63</y>
    </pos>
    <children>
      <id>6530d4a4-4e48-4884-bca9-82460c8c1edc</id>
      <ty>SelectorNode</ty>
      <pos>
        <x>-5</x>
        <y>143</y>
      </pos>
      <children>
        <id>be495137-09e7-4135-89b1-c1fd912b3ec4</id>
        <ty>ConditionNode</ty>
        <pos>
          <x>-50</x>
          <y>190</y>
        </pos>
        <children>
          <id>757fc3b1-61be-4df2-b5bc-821271f88b5a</id>
          <ty>ActionNode</ty>
          <pos>
            <x>-60</x>
            <y>250</y>
          </pos>
          <code>
local parm = {
  body = {},    -- request body
  timeout = &#34;10s&#34;,
  headers = {},
}

local url = &#34;http://127.0.0.1:6666/login/guest&#34;
local http = require(&#34;http&#34;)

function execute()

  -- http post request
  res, errmsg = http.post(url, parm)
  if errmsg == nil then
    body = json.decode(res[&#34;body&#34;])
    merge(meta, body.Body)
  end
end
</code>
          <alias>login/guest</alias>
        </children>
        <code>

-- Write expression to return true or false
function execute()

    return meta.Token == &#34;&#34;

end
</code>
        <alias></alias>
      </children>
      <children>
        <id>1adab96f-f884-4171-bd60-30c04aeaf2f5</id>
        <ty>ConditionNode</ty>
        <pos>
          <x>86</x>
          <y>190</y>
        </pos>
        <children>
          <id>811ecafa-58cc-4475-a265-4384448e3df6</id>
          <ty>ActionNode</ty>
          <pos>
            <x>76</x>
            <y>240</y>
          </pos>
          <children>
            <id>2f347fe4-9110-45c1-b519-e17458e94176</id>
            <ty>SequenceNode</ty>
            <pos>
              <x>66</x>
              <y>338</y>
            </pos>
            <children>
              <id>53ea97df-def8-4359-b09c-276498c737ef</id>
              <ty>ActionNode</ty>
              <pos>
                <x>0</x>
                <y>401</y>
              </pos>
              <code>
local parm = {
  body = {
      Token = meta.Token
  },    -- request body
  timeout = &#34;10s&#34;,
  headers = {},
}

local url = &#34;http://127.0.0.1:6666/base/hero.info&#34;
local http = require(&#34;http&#34;)

function execute()

  -- http post request
  res, errmsg = http.post(url, parm)
  if errmsg == nil then
    body = json.decode(res[&#34;body&#34;])
    merge(meta, body.Body)
  end

end
</code>
              <alias>base/hero.info</alias>
            </children>
            <children>
              <id>64be51e8-4e7e-41ef-b8c5-2a436bf92885</id>
              <ty>WaitNode</ty>
              <pos>
                <x>86</x>
                <y>401</y>
              </pos>
              <wait>100</wait>
            </children>
            <children>
              <id>cf82ce82-a0b8-46ea-9141-3f197d686e99</id>
              <ty>LoopNode</ty>
              <pos>
                <x>167</x>
                <y>401</y>
              </pos>
              <children>
                <id>cb519324-d8fb-4519-ab3d-48c14dbe4b35</id>
                <ty>ActionNode</ty>
                <pos>
                  <x>177</x>
                  <y>476</y>
                </pos>
                <code>
local parm = {
  body = {
      Token = meta.Token,
      HeroID = &#34;joy&#34;
  },    -- request body
  timeout = &#34;10s&#34;,
  headers = {},
}

local url = &#34;http://127.0.0.1:6666/base/hero.lvup&#34;
local http = require(&#34;http&#34;)

function execute()

  -- http post request
  res, errmsg = http.post(url, parm)
  if errmsg == nil then
    body = json.decode(res[&#34;body&#34;])
    merge(meta, body.Body)
  end

end
</code>
                <alias>base/hero.lvup</alias>
              </children>
              <loop>5</loop>
            </children>
          </children>
          <code>
local parm = {
  body = {
      Token = meta.Token
  },    -- request body
  timeout = &#34;10s&#34;,
  headers = {},
}

local url = &#34;http://127.0.0.1:6666/base/acc.info&#34;
local http = require(&#34;http&#34;)

function execute()

  -- http post request
  res, errmsg = http.post(url, parm)
  if errmsg == nil then
    body = json.decode(res[&#34;body&#34;])
    merge(meta, body.Body)
  end

end
</code>
          <alias>base/acc.info</alias>
        </children>
        <code>

-- Write expression to return true or false
function execute()

    return meta.Token ~= &#34;&#34;

end
</code>
        <alias></alias>
      </children>
    </children>
    <loop>5</loop>
  </children>
</behavior>
`

func TestLoadTree(t *testing.T) {

	_, err := Load([]byte(compose))
	assert.Equal(t, err, nil)

}

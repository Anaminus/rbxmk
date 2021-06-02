-- Output a CSV list of transactions for the user logged into Studio.
--
-- List possible transaction types:
--     rbxmk run fetch-transactions.rbxmk.lua
--
-- List sales:
--     rbxmk run fetch-transactions.rbxmk.lua sale

local txnTypes = {
	adspend               = true,
	affiliatesale         = true,
	csadjustment          = true,
	currencypurchase      = true,
	devex                 = true,
	engagementpayout      = true,
	groupengagementpayout = true,
	grouppayout           = true,
	individualtogroup     = true,
	premiumstipend        = true,
	purchase              = true,
	sale                  = true,
	traderobux            = true,
}
local txnType = string.lower(tostring((...) or ""))
if not txnTypes[txnType] then
	local s = {}
	for k in pairs(txnTypes) do
		table.insert(s, k)
	end
	table.sort(s)
	print("first argument must be one of the following transaction types:")
	for i, k in ipairs(s) do
		print("\t"..k)
	end
	return
end

local cookies = rbxmk.cookiesFrom("studio")

local CurrentUser do
	local userURL = "https://users.roblox.com/v1/users/authenticated"
	local resp = http.request({
		URL = userURL,
		Method = "GET",
		ResponseFormat = "json",
		Cookies = cookies,
	}):Resolve()
	if not resp.Success then
		print("failed to get user info:", resp.StatusMessage)
		return
	end
	CurrentUser = {
		ID = resp.Body.id,
		name = resp.Body.name,
		displayName = resp.Body.displayName,
	}
end

local function getPage(next)
	local txnURL = "https://economy.roblox.com/v2/users/%d/transactions?transactionType=%s&limit=100&cursor=%s"
	local resp = http.request({
		URL = string.format(txnURL, CurrentUser.ID, txnType, next or ""),
		Method = "GET",
		ResponseFormat = "json",
		Cookies = cookies,
	}):Resolve()
	if not resp.Success then
		return nil
	end
	local rows = {}
	for i, entry in ipairs(resp.Body.data) do
		local row = {
			tostring(entry.id),
			tostring(entry.transactionType),
			tostring(entry.created),
			tostring(entry.isPending),
			tostring(entry.agent.id),
			tostring(entry.agent.type),
			tostring(entry.agent.name),
			entry.details and tostring(entry.details.id) or "",
			entry.details and tostring(entry.details.name) or "",
			entry.details and tostring(entry.details.type) or "",
			tostring(entry.currency.amount),
			tostring(entry.currency.type),
		}
		table.insert(rows, row)
	end
	print(rbxmk.encodeFormat("csv", rows):sub(1,-2))
	return resp.Body.nextPageCursor
end

local rows = {{
	"id",
	"transactionType",
	"created",
	"isPending",
	"agent.id",
	"agent.type",
	"agent.name",
	"details.id",
	"details.name",
	"details.type",
	"currency.amount",
	"currency.type",
}}
print(rbxmk.encodeFormat("csv", rows):sub(1,-2))
local next
repeat
	next = getPage(next)
	if not next then
		break
	end
until not next

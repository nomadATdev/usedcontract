// Copyright (c) 2026 UsedContract contributors
// SPDX-License-Identifier: BUSL-1.1
using System.Net;
using System.Text.Json;

using var handler = new CaptureHandler();
using var http = new HttpClient(handler) { BaseAddress = new Uri("http://collector.test/") };
var agent = new UsedContract.Agent(http, "checkout", "1.0");
await agent.CompareAsync("GET", "/account", "$.status", "eq", "ACTIVE", true);
using (var doc = JsonDocument.Parse(handler.Body!))
{
    var e = doc.RootElement;
    Require(e.GetProperty("operator").GetString() == "eq", "Canonical operator field missing");
    Require(!e.TryGetProperty("@operator", out _), "C# escaping leaked into JSON");
    Require(e.GetProperty("operand").GetString() == "ACTIVE", "Operand changed");
    Require(e.GetProperty("result").GetBoolean(), "Comparison result changed");
    Require(e.GetProperty("consumer").GetProperty("language").GetString() == "dotnet", "Consumer language missing");
    Require(handler.Endpoint == "/v1/events", "Collector endpoint changed");
}
await agent.ReadAsync("GET", "/account", "$.status");
using (var doc = JsonDocument.Parse(handler.Body!))
    Require(doc.RootElement.GetProperty("kind").GetString() == "read", "Read event changed");
await agent.ControlAsync("GET", "/account", "$.status", false);
using (var doc = JsonDocument.Parse(handler.Body!))
    Require(!doc.RootElement.GetProperty("result").GetBoolean(), "False control result lost");
handler.Status = HttpStatusCode.ServiceUnavailable;
bool rejected = false;
try { await agent.ReadAsync("GET", "/account", "$.status"); }
catch (HttpRequestException) { rejected = true; }
Require(rejected, "Collector failure was swallowed");
Console.WriteLine(".NET adapter: compare/read/control serialization and HTTP failure checks passed.");

static void Require(bool value, string message)
{
    if (!value) throw new InvalidOperationException(message);
}

sealed class CaptureHandler : HttpMessageHandler
{
    public string? Body { get; private set; }
    public string? Endpoint { get; private set; }
    public HttpStatusCode Status { get; set; } = HttpStatusCode.Accepted;
    protected override async Task<HttpResponseMessage> SendAsync(HttpRequestMessage request, CancellationToken ct)
    {
        Body = await request.Content!.ReadAsStringAsync(ct);
        Endpoint = request.RequestUri!.AbsolutePath;
        return new HttpResponseMessage(Status);
    }
}

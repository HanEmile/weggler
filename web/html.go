package web

const indexHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width,initial-scale=1">
<title>weggler</title>
<style>
@import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@300;400;500;600;700&family=Inter:wght@400;500;600&display=swap');
:root{
  --bg:#050608;--bg2:#0a0c10;--bg3:#111418;--bg4:#181c22;--bg5:#1e242e;
  --b1:#2a3548;--b2:#364560;--b3:#425572;
  --tx:#dce6f0;--txd:#7090aa;--txm:#3a4e62;
  --ac:#28e070;--ac2:#20c8e0;--ac3:#e0a020;--ac4:#b060f8;
  --red:#e84040;--run:#7080f0;
  --mono:'JetBrains Mono',monospace;--sans:'Inter',sans-serif;
  --r:3px;
}
*{box-sizing:border-box;margin:0;padding:0;}
html,body{height:100%;overflow:hidden;}
body{background:var(--bg);color:var(--tx);font-family:var(--sans);font-size:12px;display:flex;flex-direction:column;}

/* ── Layout ── */
.shell{display:flex;flex-direction:column;height:100vh;}
header{height:34px;display:flex;align-items:stretch;background:var(--bg2);border-bottom:2px solid var(--b2);flex-shrink:0;padding-left:10px;}
.logo{font-family:var(--mono);font-weight:700;font-size:12px;color:var(--ac);display:flex;align-items:center;margin-right:14px;letter-spacing:-.3px;}
.logo em{color:var(--txm);font-style:normal;}
nav{display:flex;}
.tab{padding:0 10px;display:flex;align-items:center;font-family:var(--mono);font-size:10px;font-weight:500;color:var(--txd);cursor:pointer;letter-spacing:.4px;text-transform:uppercase;border-bottom:2px solid transparent;margin-bottom:-2px;transition:color .1s,border-color .1s;user-select:none;gap:5px;}
.tab:hover{color:var(--tx);background:var(--bg3);}
.tab.active{color:var(--ac);border-bottom-color:var(--ac);}
.tab .tc{font-size:9px;padding:0 4px;border-radius:2px;line-height:1.6;}
.hright{margin-left:auto;display:flex;align-items:center;gap:8px;padding:0 10px;}
.pulse{width:5px;height:5px;border-radius:50%;background:var(--ac);box-shadow:0 0 5px var(--ac);animation:pa 2s ease-in-out infinite;}
@keyframes pa{0%,100%{opacity:1;}50%{opacity:.25;box-shadow:none;}}
.ver{font-family:var(--mono);font-size:9px;color:var(--txm);}
.main{flex:1;min-height:0;display:flex;}
.page{display:none;flex:1;min-height:0;overflow-y:auto;padding:10px 12px;flex-direction:column;gap:8px;}
.page.active{display:flex;}
/* Direct children of .page must not shrink — they push the page taller so it scrolls */
.page > *{flex-shrink:0;}
/* findings page manages its own full-viewport layout — must not be flex-shrunk */
#page-findings{overflow:hidden;}
#page-findings > *{flex-shrink:1;min-height:0;}

/* ── Panels ── */
.panel{background:var(--bg2);border:1px solid var(--b1);border-radius:var(--r);overflow:hidden;}
.ph{display:flex;align-items:center;justify-content:space-between;padding:6px 10px;border-bottom:1px solid var(--b1);background:var(--bg3);}
.pt{font-family:var(--mono);font-size:9px;font-weight:700;text-transform:uppercase;letter-spacing:1px;color:var(--txd);}
.pb{padding:8px 10px;}

/* ── Buttons ── */
.btn{display:inline-flex;align-items:center;gap:4px;padding:3px 8px;font-family:var(--mono);font-size:10px;font-weight:500;border-radius:var(--r);border:1px solid;cursor:pointer;transition:background .1s,color .1s;text-transform:uppercase;letter-spacing:.3px;background:transparent;white-space:nowrap;line-height:1.5;}
.btn-p{color:var(--ac);border-color:var(--ac);}    .btn-p:hover{background:rgba(40,224,112,.1);}
.btn-s{color:var(--txd);border-color:var(--b2);}   .btn-s:hover{color:var(--tx);border-color:var(--b3);background:var(--bg4);}
.btn-d{color:var(--red);border-color:var(--red);}  .btn-d:hover{background:rgba(232,64,64,.1);}
.btn-a{color:var(--ac4);border-color:var(--ac4);}  .btn-a:hover{background:rgba(176,96,248,.1);}
.btn-2{color:var(--ac2);border-color:var(--ac2);}  .btn-2:hover{background:rgba(32,200,224,.1);}
.btn:disabled{opacity:.3;cursor:not-allowed;}
.bgroup{display:flex;gap:4px;}

/* ── Forms ── */
.fg{display:flex;flex-direction:column;gap:3px;}
.fl{font-family:var(--mono);font-size:9px;font-weight:700;text-transform:uppercase;letter-spacing:.7px;color:var(--txd);}
.fi,.fs,.fta{background:var(--bg3);border:1px solid var(--b1);border-radius:var(--r);color:var(--tx);font-family:var(--sans);font-size:12px;padding:5px 7px;outline:none;transition:border-color .12s;width:100%;}
.fi:focus,.fs:focus,.fta:focus{border-color:var(--ac);}
.fta{font-family:var(--mono);font-size:11px;resize:vertical;min-height:80px;}
.fta.tall{min-height:120px;}
.fs{appearance:none;cursor:pointer;}
.fgrid{display:grid;gap:9px;}
.frow{display:grid;grid-template-columns:1fr 1fr;gap:9px;}
.frow3{display:grid;grid-template-columns:1fr 1fr 1fr;gap:9px;}
.rgroup{display:flex;gap:5px;flex-wrap:wrap;}
.ropt{display:flex;align-items:center;gap:3px;padding:3px 8px;border:1px solid var(--b1);border-radius:var(--r);cursor:pointer;font-family:var(--mono);font-size:10px;color:var(--txd);}
.ropt input{accent-color:var(--ac);}
.ropt.sel{border-color:var(--ac);color:var(--ac);background:rgba(40,224,112,.05);}

/* ── Item list ── */
.ilist{display:flex;flex-direction:column;}
.item{display:flex;align-items:center;gap:7px;padding:6px 10px;border-bottom:1px solid var(--b1);transition:background .08s;cursor:pointer;}
.item:last-child{border-bottom:none;}
.item:hover{background:var(--bg3);}
.iico{width:20px;height:20px;display:flex;align-items:center;justify-content:center;border-radius:var(--r);font-size:9px;flex-shrink:0;font-family:var(--mono);font-weight:700;}
.ibody{flex:1;min-width:0;}
.iname{font-weight:500;font-size:12px;white-space:nowrap;overflow:hidden;text-overflow:ellipsis;}
.imeta{font-family:var(--mono);font-size:9px;color:var(--txd);margin-top:1px;white-space:nowrap;overflow:hidden;text-overflow:ellipsis;}
.iid{font-family:var(--mono);font-size:8px;color:var(--txm);margin-top:1px;}
.iact{display:flex;gap:3px;flex-shrink:0;align-items:center;}

/* ── Badges ── */
.bdg{display:inline-flex;align-items:center;padding:1px 5px;border-radius:2px;font-family:var(--mono);font-size:9px;font-weight:600;text-transform:uppercase;border:1px solid;}
.bdg-cpp{color:var(--ac2);border-color:var(--ac2);background:rgba(32,200,224,.07);}
.bdg-c{color:var(--ac3);border-color:var(--ac3);background:rgba(224,160,32,.07);}
.bdg-both{color:var(--ac4);border-color:var(--ac4);background:rgba(176,96,248,.07);}
.bdg-done{color:var(--ac);border-color:var(--ac);background:rgba(40,224,112,.07);}
.bdg-running{color:var(--run);border-color:var(--run);background:rgba(112,128,240,.07);animation:blink 1.3s ease-in-out infinite;}
.bdg-queued{color:var(--ac3);border-color:var(--ac3);background:rgba(224,160,32,.07);}
.bdg-failed{color:var(--red);border-color:var(--red);background:rgba(232,64,64,.07);}
.bdg-pending,.bdg-canceled{color:var(--txd);border-color:var(--b2);}
@keyframes blink{0%,100%{opacity:1;}50%{opacity:.35;}}

/* confidence pill */
.conf{display:inline-flex;align-items:center;padding:1px 5px;border-radius:2px;font-family:var(--mono);font-size:9px;border:1px solid;}
.conf-hi{color:#28e070;border-color:#28e070;background:rgba(40,224,112,.07);}
.conf-mid{color:#e0a020;border-color:#e0a020;background:rgba(224,160,32,.07);}
.conf-lo{color:#e84040;border-color:#e84040;background:rgba(232,64,64,.07);}

/* ── Stats ── */
.sgrid{display:grid;grid-template-columns:repeat(6,1fr);gap:6px;flex-shrink:0;}
.sc{background:var(--bg2);border:1px solid var(--b1);border-radius:var(--r);padding:8px 10px;}
.sv{font-family:var(--mono);font-size:20px;font-weight:700;line-height:1;margin-bottom:2px;}
.sl{font-family:var(--mono);font-size:9px;text-transform:uppercase;letter-spacing:.8px;color:var(--txd);}

/* ── Metrics bar ── */
.mbar{display:flex;align-items:center;gap:8px;padding:5px 10px;background:var(--bg2);border:1px solid var(--b1);border-radius:var(--r);font-family:var(--mono);font-size:10px;color:var(--txd);}
.mbar b{color:var(--tx);}
.mbar-sep{width:1px;height:12px;background:var(--b2);}
.mbar-right{margin-left:auto;display:flex;align-items:center;gap:6px;}

/* ── Terminal ── */
.term{background:#030405;border:1px solid var(--b1);border-radius:var(--r);overflow:hidden;display:flex;flex-direction:column;}
.tbar{display:flex;align-items:center;gap:5px;padding:4px 8px;background:var(--bg2);border-bottom:1px solid var(--b1);flex-shrink:0;}
.tdot{width:7px;height:7px;border-radius:50%;}
.ttitle{font-family:var(--mono);font-size:9px;color:var(--txd);flex:1;margin-left:3px;}
.tout{font-family:var(--mono);font-size:10.5px;line-height:1.6;padding:8px 10px;color:#5a7080;flex:1;overflow-y:auto;min-height:120px;max-height:320px;}
.tline{white-space:pre-wrap;word-break:break-all;min-height:1em;}
.tline:empty{min-height:1em;}
/* ── Findings table ── */
.ftable{width:100%;border-collapse:collapse;}
.ftable th{font-family:var(--mono);font-size:9px;font-weight:700;text-transform:uppercase;letter-spacing:.7px;color:var(--txd);padding:5px 9px;text-align:left;border-bottom:1px solid var(--b2);white-space:nowrap;background:var(--bg3);position:sticky;top:0;z-index:1;}
.ftable td{padding:4px 9px;border-bottom:1px solid var(--b1);vertical-align:middle;}
.ftable tr:last-child td{border-bottom:none;}
.ftable tr:hover td{background:var(--bg3);cursor:pointer;}
.fc{font-family:var(--mono);font-size:10px;color:var(--ac2);max-width:240px;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;display:block;}
.fn{font-family:var(--mono);font-size:10px;color:var(--ac3);white-space:nowrap;}
.fs2{font-family:var(--mono);font-size:10px;color:var(--txd);max-width:260px;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;display:block;}

/* ── Graph ── */
.findings-page{display:flex;flex-direction:column;height:calc(100vh - 34px);}
.fbar2{display:flex;align-items:center;gap:6px;padding:5px 8px;background:var(--bg2);border-bottom:2px solid var(--b2);flex-shrink:0;}
.fbar2 .fi{max-width:200px;font-size:10px;padding:3px 6px;}
.vtab2{padding:4px 10px;font-family:var(--mono);font-size:9px;font-weight:700;text-transform:uppercase;letter-spacing:.4px;color:var(--txd);cursor:pointer;border-bottom:2px solid transparent;margin-bottom:-2px;transition:all .1s;}
.vtab2.active{color:var(--ac);border-bottom-color:var(--ac);}
.vtab2:hover{color:var(--tx);}
.fcount{font-family:var(--mono);font-size:9px;color:var(--txd);}
.graph-wrap{flex:1;position:relative;background:#030405;overflow:hidden;}
#graphCanvas{display:block;width:100%;height:100%;}
.graph-legend{position:absolute;top:6px;right:8px;display:flex;flex-direction:column;gap:3px;pointer-events:none;}
.gl-item{display:flex;align-items:center;gap:4px;font-family:var(--mono);font-size:9px;color:var(--txd);}
.gl-dot{width:6px;height:6px;border-radius:50%;flex-shrink:0;}
.graph-tip{position:absolute;pointer-events:none;background:rgba(8,10,14,.97);border:1px solid var(--b2);border-radius:var(--r);padding:5px 8px;font-family:var(--mono);font-size:10px;color:var(--tx);display:none;max-width:300px;word-break:break-all;line-height:1.55;z-index:10;}
.graph-ctrl{position:absolute;bottom:6px;left:8px;display:flex;gap:4px;}
.table-wrap{overflow:auto;flex:1;min-height:0;}

/* ── Scripts tree (details/summary) ── */
.stree{display:flex;flex-direction:column;}
.stree-filter{display:flex;gap:5px;align-items:center;padding:6px 10px;border-bottom:1px solid var(--b1);background:var(--bg3);}
.stree-filter .fi{font-size:11px;padding:3px 7px;}
details.gd{border-bottom:1px solid var(--b1);}
details.gd:last-child{border-bottom:none;}
details.gd>summary{display:flex;align-items:center;gap:6px;padding:5px 10px;cursor:pointer;list-style:none;background:var(--bg3);border-left:2px solid transparent;transition:background .08s;}
details.gd>summary::-webkit-details-marker{display:none;}
details.gd>summary:hover{background:var(--bg4);}
details.gd[open]>summary{border-left-color:var(--ac4);background:var(--bg4);}
.sum-arr{font-family:var(--mono);font-size:10px;color:var(--txd);transition:transform .15s;flex-shrink:0;width:10px;}
details.gd[open]>.sum-arr-wrap>.sum-arr{display:none;}
.sum-lbl{font-weight:500;font-size:12px;flex:1;min-width:0;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;}
.sum-cnt{font-family:var(--mono);font-size:9px;color:var(--txd);flex-shrink:0;}
.gd-body{padding-left:14px;border-left:1px solid var(--b1);}
.gd-body details.gd>summary{padding-left:14px;}
.gd-body .gd-body details.gd>summary{padding-left:24px;}
.sc-row{display:flex;align-items:center;gap:6px;padding:5px 10px;border-bottom:1px solid var(--b1);transition:background .08s;}
.sc-row:last-child{border-bottom:none;}
.sc-row:hover{background:var(--bg3);}
.sc-row-body{flex:1;min-width:0;}
.ungrouped-hdr{font-family:var(--mono);font-size:9px;font-weight:700;text-transform:uppercase;letter-spacing:1px;color:var(--txm);padding:5px 10px;background:var(--bg3);border-bottom:1px solid var(--b1);}

/* ── Check list ── */
.check-list{display:flex;flex-direction:column;gap:2px;max-height:160px;overflow-y:auto;background:var(--bg3);border:1px solid var(--b1);border-radius:var(--r);padding:4px 6px;}
.ci{display:flex;align-items:center;gap:5px;padding:2px 3px;border-radius:2px;cursor:pointer;}
.ci:hover{background:var(--bg4);}
.ci input{accent-color:var(--ac);width:11px;height:11px;}
.ci label{font-family:var(--mono);font-size:10px;color:var(--tx);cursor:pointer;flex:1;}
.ci-sub{color:var(--txd);font-size:9px;}

/* ── Split ── */
.split2{display:grid;grid-template-columns:1fr 1fr;gap:8px;}

/* ── Modal ── */
.overlay{display:none;position:fixed;inset:0;background:rgba(0,0,0,.8);backdrop-filter:blur(3px);z-index:100;align-items:center;justify-content:center;}
.overlay.open{display:flex;}
.modal{background:var(--bg2);border:1px solid var(--b2);border-radius:4px;width:560px;max-width:calc(100vw - 20px);max-height:calc(100vh - 40px);overflow-y:auto;box-shadow:0 20px 60px rgba(0,0,0,.8);}
.modal-lg{width:680px;}
.mh{display:flex;align-items:center;justify-content:space-between;padding:8px 12px;border-bottom:1px solid var(--b1);}
.mt{font-family:var(--mono);font-size:10px;font-weight:700;text-transform:uppercase;letter-spacing:.4px;color:var(--tx);}
.mx{background:none;border:none;color:var(--txd);cursor:pointer;font-size:14px;line-height:1;padding:1px 4px;border-radius:2px;}
.mx:hover{color:var(--tx);background:var(--bg4);}
.mb{padding:12px;}
.mf{display:flex;justify-content:flex-end;gap:5px;padding:8px 12px;border-top:1px solid var(--b1);}

/* ── Misc ── */
.empty{text-align:center;padding:22px 14px;color:var(--txm);font-family:var(--mono);font-size:10px;}
.empty-ico{font-size:20px;margin-bottom:6px;opacity:.2;}
.cmd-prev{font-family:var(--mono);font-size:10px;color:var(--txd);background:var(--bg3);border:1px solid var(--b1);border-radius:var(--r);padding:6px 9px;word-break:break-all;line-height:1.7;}
.cmd-prev .cb{color:var(--ac);} .cmd-prev .cf{color:var(--ac3);} .cmd-prev .cp{color:var(--ac2);} .cmd-prev .cx{color:var(--tx);}
.spin{display:inline-block;width:9px;height:9px;border:2px solid var(--b2);border-top-color:var(--ac);border-radius:50%;animation:spin .6s linear infinite;vertical-align:middle;}
@keyframes spin{to{transform:rotate(360deg);}}
.snip-modal{background:#030405;border:1px solid var(--b1);border-radius:var(--r);padding:8px 10px;font-family:var(--mono);font-size:11px;line-height:1.6;color:var(--tx);max-height:260px;overflow-y:auto;}
.lnk{color:var(--ac2);text-decoration:none;font-family:var(--mono);font-size:11px;cursor:pointer;}
.lnk:hover{text-decoration:underline;}
.monoval{font-family:var(--mono);font-size:11px;}
.tago{font-family:var(--mono);font-size:9px;color:var(--txm);}
.tagf{font-family:var(--mono);font-size:9px;color:var(--txd);background:var(--bg3);border:1px solid var(--b1);padding:1px 5px;border-radius:2px;}
.size-info{font-family:var(--mono);font-size:9px;color:var(--txd);}
.est-time{font-family:var(--mono);font-size:9px;color:var(--ac3);}
::-webkit-scrollbar{width:3px;height:3px;}::-webkit-scrollbar-track{background:transparent;}::-webkit-scrollbar-thumb{background:var(--b2);border-radius:2px;}
</style>
</head>
<body>
<div class="shell">
<header>
  <div class="logo">weggler<em>/</em></div>
  <nav>
    <div class="tab active" onclick="showPage('dashboard')">Dashboard</div>
    <div class="tab" onclick="showPage('jobs')">Jobs <span class="tc bdg-running" id="tabRunning" style="display:none"></span></div>
    <div class="tab" onclick="showPage('findings')">Findings</div>
    <div class="tab" onclick="showPage('scripts')">Scripts</div>
    <div class="tab" onclick="showPage('sources')">Sources</div>
  </nav>
  <div class="hright"><div class="pulse"></div><button class="btn btn-s" style="font-size:9px;padding:2px 7px" onclick="reloadData()" id="reloadBtn">↺ Reload</button><div class="ver">weggler</div></div>
</header>
<div class="main">

<!-- ══ DASHBOARD ══ -->
<div class="page active" id="page-dashboard">
  <div class="sgrid">
    <div class="sc"><div class="sv" id="stat-jobs" style="color:var(--ac2)">–</div><div class="sl">Jobs</div></div>
    <div class="sc"><div class="sv" id="stat-running" style="color:var(--run)">–</div><div class="sl">Running</div></div>
    <div class="sc"><div class="sv" id="stat-queued" style="color:var(--ac3)">–</div><div class="sl">Queued</div></div>
    <div class="sc"><div class="sv" id="stat-findings" style="color:var(--red)">–</div><div class="sl">Findings</div></div>
    <div class="sc"><div class="sv" id="stat-scripts" style="color:var(--ac3)">–</div><div class="sl">Scripts</div></div>
    <div class="sc"><div class="sv" id="stat-sources" style="color:var(--ac)">–</div><div class="sl">Sources</div></div>
  </div>
  <div id="metricsBar" class="mbar">
    <span>max-jobs: <b id="m-max">–</b></span><div class="mbar-sep"></div>
    <span>running: <b id="m-run">–</b></span><div class="mbar-sep"></div>
    <span>queued: <b id="m-que">–</b></span><div class="mbar-sep"></div>
    <span>completed: <b id="m-comp">–</b></span><div class="mbar-sep"></div>
    <span>total findings: <b id="m-find">–</b></span><div class="mbar-sep"></div>
    <span>queue ETA: <b id="m-eta" style="color:var(--ac3)">–</b></span>
    <div class="mbar-right">
      <span style="font-family:var(--mono);font-size:9px;color:var(--txd)">max-jobs:</span>
      <input id="maxJobsInput" class="fi" type="number" min="1" max="64" style="width:52px;padding:2px 5px;font-size:10px" onchange="setMaxJobs()">
    </div>
  </div>
  <div class="panel">
    <div class="ph">
      <div class="pt">Recent Jobs</div>
      <div class="bgroup">
        <button class="btn btn-s" onclick="openScriptRunModal()">▶ Script</button><button class="btn btn-a" onclick="openGroupRunModal()">▶ Group</button>
        <button class="btn btn-p" onclick="openNewJobModal()">+ Job</button>
      </div>
    </div>
    <div class="ilist" id="dashJobList"><div class="empty"><div class="empty-ico">◈</div>No jobs</div></div>
  </div>
</div>

<!-- ══ JOBS ══ -->
<div class="page" id="page-jobs">
  <div class="panel">
    <div class="ph">
      <div class="pt">Jobs</div>
      <div class="bgroup">
        <button class="btn btn-s" onclick="openScriptRunModal()">▶ Script</button><button class="btn btn-a" onclick="openGroupRunModal()">▶ Group</button>
        <button class="btn btn-p" onclick="openNewJobModal()">+ Job</button>
      </div>
    </div>
    <div class="ilist" id="jobList"><div class="empty"><div class="empty-ico">◈</div>No jobs</div></div>
  </div>
</div>

<!-- ══ JOB DETAIL ══ -->
<div class="page" id="page-job-detail">
  <div style="display:flex;align-items:center;gap:7px;flex-shrink:0">
    <button class="btn btn-s" onclick="showPage('jobs')">← Back</button>
    <span id="jd-title" style="font-weight:600;font-size:12px;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;max-width:400px"></span>
    <span id="jd-badge"></span>
    <div style="margin-left:auto;display:flex;gap:4px">
      <button class="btn btn-d" id="jd-cancel-btn" style="display:none" onclick="cancelCurrentJob()">Cancel</button>
      <button class="btn btn-p" id="jd-restart-btn" style="display:none" onclick="restartCurrentJob()">↺ Retry</button>
      <button class="btn btn-s" id="jd-findings-btn" style="display:none" onclick="openJobFindings()">Findings →</button>
    </div>
  </div>
  <div class="split2">
    <div class="panel">
      <div class="ph"><div class="pt">Details</div></div>
      <div class="pb" style="display:grid;gap:7px">
        <div class="frow">
          <div class="fg"><div class="fl">Script</div><a class="lnk" id="jd-script-lnk" href="#" onclick="return jdNavScript(event)"></a>
            <div id="jd-script-desc" style="font-size:10px;color:var(--txd);margin-top:2px"></div>
          </div>
          <div class="fg"><div class="fl">Language / Confidence</div><div id="jd-script-meta" style="display:flex;gap:4px;align-items:center;margin-top:2px"></div></div>
        </div>
        <div class="fg"><div class="fl">Pattern</div><div id="jd-script-pattern" class="cmd-prev" style="font-size:10px;max-height:80px;overflow-y:auto"></div></div>
        <div class="fg"><div class="fl">Source</div><a class="lnk" id="jd-source-lnk" href="#" onclick="return jdNavSource(event)"></a></div>
        <div class="frow3">
          <div class="fg"><div class="fl">Started</div><div id="jd-started" class="tago"></div></div>
          <div class="fg"><div class="fl">Duration</div><div id="jd-duration" class="tago"></div></div>
          <div class="fg"><div class="fl">Findings</div><div id="jd-fcount" style="font-family:var(--mono);font-size:11px;color:var(--red)">–</div></div>
        </div>
        <div class="fg"><div class="fl">Description</div><div id="jd-desc" style="font-size:11px;color:var(--txd)"></div></div>
        <div class="fg"><div class="fl">Source Size</div><div id="jd-srcsize" class="size-info"></div></div>
      </div>
    </div>
    <div class="fg">
      <div class="fl">Command</div>
      <div class="cmd-prev" id="jd-cmd"></div>
    </div>
  </div>
  <div class="term">
    <div class="tbar">
      <div class="tdot" style="background:#b83030"></div>
      <div class="tdot" style="background:#b89020"></div>
      <div class="tdot" style="background:#20a040"></div>
      <div class="ttitle" id="jd-ttitle">output</div>
      <div id="jd-spin" style="display:none"><div class="spin"></div></div>
    </div>
    <div class="tout" id="jd-out"></div>
  </div>
</div>

<!-- ══ FINDINGS ══ -->
<div class="page" id="page-findings" style="padding:0;overflow:hidden;">
  <div class="findings-page">
    <div class="fbar2">
      <div class="vtab2 active" onclick="switchVTab('graph')">Graph</div>
      <div class="vtab2" onclick="switchVTab('table')">Table</div>
      <div style="width:1px;height:14px;background:var(--b2);margin:0 2px"></div>
      <select class="fs" id="findings-job-sel" style="max-width:200px;font-size:10px;padding:3px 6px;" onchange="loadFindings()"><option value="">— all jobs —</option></select>
      <input class="fi" id="ftFilter" style="max-width:180px;font-size:10px;padding:3px 6px;" placeholder="filter…" oninput="renderFindingTable()">
      <span class="fcount" id="ftCount"></span>
      <div style="margin-left:auto;display:flex;gap:4px;align-items:center">
        <button class="btn btn-s" style="font-size:9px;padding:2px 6px" onclick="resetGraphView()">Reset</button>
      </div>
    </div>
    <!-- graph view -->
    <div id="vp-graph" style="flex:1;display:flex;flex-direction:column;overflow:hidden">
      <div class="graph-wrap" id="graphWrap">
        <canvas id="graphCanvas"></canvas>
        <div class="graph-legend" id="graphLegend"></div>
        <div class="graph-tip" id="graphTip"></div>
        <div class="graph-ctrl">
          <div style="font-family:var(--mono);font-size:9px;color:var(--txd);background:rgba(5,6,8,.85);padding:3px 6px;border-radius:2px;border:1px solid var(--b1)">
            scroll=zoom · drag=pan · click finding=detail
          </div>
        </div>
      </div>
    </div>
    <!-- table view -->
    <div id="vp-table" style="flex:1;display:none;flex-direction:column;min-height:0;">
      <div class="table-wrap">
        <table class="ftable">
          <thead><tr><th>#</th><th>File</th><th>Line</th><th>Snippet</th><th>Script</th><th>Source</th><th>Conf</th><th>Found</th></tr></thead>
          <tbody id="ftBody"></tbody>
        </table>
      </div>
    </div>
  </div>
</div>

<!-- ══ SCRIPTS ══ -->
<div class="page" id="page-scripts">
  <div class="panel">
    <div class="ph">
      <div class="pt">Scripts &amp; Groups</div>
      <div class="bgroup">
        <button class="btn btn-a" onclick="openGroupModal()">+ Group</button>
        <button class="btn btn-p" onclick="openScriptModal()">+ Script</button>
      </div>
    </div>
    <div class="stree-filter">
      <input class="fi" id="scriptSearch" placeholder="Search scripts…" oninput="renderScriptTree()">
      <div class="rgroup" id="langFilter">
        <label class="ropt sel" id="lf-all"><input type="radio" name="lf" value="" checked> All</label>
        <label class="ropt" id="lf-c"><input type="radio" name="lf" value="c"> C</label>
        <label class="ropt" id="lf-cpp"><input type="radio" name="lf" value="cpp"> C++</label>
        <label class="ropt" id="lf-both"><input type="radio" name="lf" value="both"> Both</label>
      </div>
      <label style="font-family:var(--mono);font-size:9px;color:var(--txd);display:flex;align-items:center;gap:4px">
        min conf: <input id="confFilter" type="number" min="0" max="100" value="0" style="width:40px" class="fi" oninput="renderScriptTree()">
      </label>
      <select class="fs" id="scriptSort" style="font-size:10px;padding:2px 6px;margin-left:auto" onchange="renderScriptTree()">
        <option value="tree">Tree order</option>
        <option value="findings">Most findings</option>
        <option value="fastest">Fastest (bytes/s)</option>
        <option value="slowest">Slowest (bytes/s)</option>
        <option value="conf">Confidence</option>
        <option value="name">Name A→Z</option>
      </select>
    </div>
    <div class="stree" id="scriptTree"></div>
  </div>
</div>

<!-- ══ SOURCES ══ -->
<div class="page" id="page-sources">
  <div class="panel">
    <div class="ph">
      <div class="pt">Sources</div>
      <button class="btn btn-p" onclick="openSourceModal()">+ Source</button>
    </div>
    <div class="ilist" id="sourceList" style="overflow-x:auto"></div>
  </div>
</div>

</div></div><!-- /main /shell -->

<!-- ── Script Modal ── -->
<div class="overlay" id="scriptModal">
  <div class="modal">
    <div class="mh"><div class="mt" id="scriptModalTitle">New Script</div><button class="mx" onclick="closeModal('scriptModal')">✕</button></div>
    <div class="mb">
      <div class="fgrid">
        <div class="frow">
          <div class="fg"><label class="fl">Name</label><input class="fi" id="sc-name"></div>
          <div class="fg"><label class="fl">Language</label>
            <div class="rgroup" id="sc-lang-group">
              <label class="ropt sel" id="sc-lang-c"><input type="radio" name="sc-lang" value="c" checked>C</label>
              <label class="ropt" id="sc-lang-cpp"><input type="radio" name="sc-lang" value="cpp">C++</label>
              <label class="ropt" id="sc-lang-both"><input type="radio" name="sc-lang" value="both">Both</label>
            </div>
          </div>
        </div>
        <div class="fg"><label class="fl">Description</label><input class="fi" id="sc-desc"></div>
        <div class="fg"><label class="fl">Pattern</label><textarea class="fta tall" id="sc-pattern"></textarea></div>
        <div class="frow">
          <div class="fg"><label class="fl">Extra Flags</label><input class="fi" id="sc-flags" placeholder="--unique"></div>
          <div class="fg">
            <label class="fl">Confidence (0–100) <span id="sc-conf-val" style="color:var(--ac3)">50</span></label>
            <input type="range" id="sc-conf" min="0" max="100" value="50" style="width:100%;accent-color:var(--ac3);margin-top:4px" oninput="document.getElementById('sc-conf-val').textContent=this.value">
          </div>
        </div>
        <div id="sc-timing-info" style="display:none;font-family:var(--mono);font-size:9px;color:var(--txd)"></div>
      </div>
    </div>
    <div class="mf"><button class="btn btn-s" onclick="closeModal('scriptModal')">Cancel</button><button class="btn btn-p" onclick="saveScript()">Save</button></div>
  </div>
</div>

<!-- ── Group Modal ── -->
<div class="overlay" id="groupModal">
  <div class="modal modal-lg">
    <div class="mh"><div class="mt" id="groupModalTitle">New Group</div><button class="mx" onclick="closeModal('groupModal')">✕</button></div>
    <div class="mb">
      <div class="fgrid">
        <div class="fg"><label class="fl">Name</label><input class="fi" id="grp-name"></div>
        <div class="fg"><label class="fl">Description</label><input class="fi" id="grp-desc"></div>
        <div class="frow">
          <div class="fg"><label class="fl">Scripts</label><div class="check-list" id="grp-scripts-list"></div></div>
          <div class="fg"><label class="fl">Child Groups</label><div class="check-list" id="grp-children-list"></div></div>
        </div>
      </div>
    </div>
    <div class="mf"><button class="btn btn-s" onclick="closeModal('groupModal')">Cancel</button><button class="btn btn-p" onclick="saveGroup()">Save</button></div>
  </div>
</div>

<!-- ── Source Modal ── -->
<div class="overlay" id="sourceModal">
  <div class="modal">
    <div class="mh"><div class="mt" id="sourceModalTitle">New Source</div><button class="mx" onclick="closeModal('sourceModal')">✕</button></div>
    <div class="mb">
      <div class="fgrid">
        <div class="frow">
          <div class="fg"><label class="fl">Name</label><input class="fi" id="src-name"></div>
        </div>
        <div class="fg"><label class="fl">Description</label><input class="fi" id="src-desc"></div>
        <div class="fg"><label class="fl">Path</label><input class="fi" id="src-path" placeholder="/path/to/source"></div>
        <div class="fg"><label class="fl">Languages</label>
          <div class="rgroup" id="src-tags-group">
            <label class="ropt" id="src-tag-c"><input type="checkbox" name="src-tag" value="c">C</label>
            <label class="ropt" id="src-tag-cpp"><input type="checkbox" name="src-tag" value="cpp">C++</label>
          </div>
        </div>
        <div id="src-stats-display" style="font-family:var(--mono);font-size:9px;color:var(--txd);display:none"></div>
      </div>
    </div>
    <div class="mf"><button class="btn btn-s" onclick="closeModal('sourceModal')">Cancel</button><button class="btn btn-p" onclick="saveSource()">Save</button></div>
  </div>
</div>

<!-- ── New Job Modal ── -->
<div class="overlay" id="jobModal">
  <div class="modal">
    <div class="mh"><div class="mt">New Job</div><button class="mx" onclick="closeModal('jobModal')">✕</button></div>
    <div class="mb">
      <div class="fgrid">
        <div class="fg"><label class="fl">Script</label><select class="fs" id="job-script-sel" onchange="updateJobPreview()"><option value="">— select script —</option></select></div>
        <div class="fg"><label class="fl">Source</label><select class="fs" id="job-source-sel" onchange="updateJobPreview()"><option value="">— select source —</option></select></div>
        <div class="fg"><label class="fl">Description (optional)</label><input class="fi" id="job-desc" placeholder="What are you investigating?"></div>
        <div id="job-cmd-preview" class="cmd-prev" style="display:none"></div>
        <div id="job-est" class="est-time" style="display:none"></div>
      </div>
    </div>
    <div class="mf"><button class="btn btn-s" onclick="closeModal('jobModal')">Cancel</button><button class="btn btn-p" id="startJobBtn" onclick="startJob()">Run</button></div>
  </div>
</div>

<!-- ── Group Run Modal ── -->
<div class="overlay" id="groupRunModal">
  <div class="modal modal-lg">
    <div class="mh"><div class="mt">Run Group</div><button class="mx" onclick="closeModal('groupRunModal')">✕</button></div>
    <div class="mb">
      <div class="fgrid">
        <div class="fg"><label class="fl">Group</label><select class="fs" id="grr-group-sel" onchange="updateGroupRunPreview()"><option value="">— select group —</option></select></div>
        <div class="fg"><label class="fl">Sources (multi-select)</label><div class="check-list" id="grr-sources-list" style="max-height:120px"></div></div>
        <div id="grr-preview" style="display:none;font-family:var(--mono);font-size:10px;color:var(--txd);background:var(--bg3);border:1px solid var(--b1);border-radius:var(--r);padding:7px 9px;line-height:1.7"></div>
      </div>
    </div>
    <div class="mf"><button class="btn btn-s" onclick="closeModal('groupRunModal')">Cancel</button><button class="btn btn-a" id="runGroupBtn" onclick="runGroup()">▶ Run All</button></div>
  </div>
</div>

<!-- ── Script Run Modal ── -->
<div class="overlay" id="scriptRunModal">
  <div class="modal modal-lg">
    <div class="mh"><div class="mt">Run Script on Sources</div><button class="mx" onclick="closeModal('scriptRunModal')">✕</button></div>
    <div class="mb">
      <div class="fgrid">
        <div class="fg"><label class="fl">Script</label><select class="fs" id="srr-script-sel" onchange="updateScriptRunPreview()"><option value="">— select script —</option></select></div>
        <div class="fg"><label class="fl">Sources</label><div class="check-list" id="srr-sources-list" style="max-height:120px"></div></div>
        <div id="srr-preview" style="display:none;font-family:var(--mono);font-size:10px;color:var(--txd);background:var(--bg3);border:1px solid var(--b1);border-radius:var(--r);padding:7px 9px;line-height:1.7"></div>
      </div>
    </div>
    <div class="mf"><button class="btn btn-s" onclick="closeModal('scriptRunModal')">Cancel</button><button class="btn btn-a" id="runScriptBtn" onclick="runScriptOnSources()">▶ Run</button></div>
  </div>
</div>

<!-- ── Finding Modal ── -->
<div class="overlay" id="findingModal">
  <div class="modal modal-lg">
    <div class="mh"><div class="mt" id="fm-title">Finding</div><button class="mx" onclick="closeModal('findingModal')">✕</button></div>
    <div class="mb" style="display:grid;gap:9px">
      <div class="frow">
        <div class="fg"><div class="fl">File</div><div id="fm-file" class="monoval" style="color:var(--ac2);word-break:break-all"></div></div>
        <div class="fg"><div class="fl">Line</div><div id="fm-line" class="monoval" style="color:var(--ac3)"></div></div>
      </div>
      <div class="frow">
        <div class="fg"><div class="fl">Script</div><a class="lnk" id="fm-script-lnk" href="#" onclick="return fmNavScript(event)"></a>
          <div id="fm-script-desc" style="font-size:10px;color:var(--txd);margin-top:2px"></div>
          <div id="fm-script-hier" style="font-family:var(--mono);font-size:9px;color:var(--txm);margin-top:2px"></div>
        </div>
        <div class="fg"><div class="fl">Source</div><a class="lnk" id="fm-source-lnk" href="#" onclick="return fmNavSource(event)"></a></div>
      </div>
      <div class="frow">
        <div class="fg"><div class="fl">Job</div><a class="lnk" id="fm-job-lnk" href="#" onclick="return fmNavJob(event)"></a></div>
        <div class="fg"><div class="fl">Found At</div><div id="fm-found-at" class="tago"></div></div>
      </div>
      <div class="fg"><div class="fl">Snippet</div><div class="snip-modal" id="fm-snippet"></div></div>
    </div>
    <div class="mf"><button class="btn btn-s" onclick="closeModal('findingModal')">Close</button></div>
  </div>
</div>

<!-- ── Timing Modal ── -->
<div class="overlay" id="timingModal">
  <div class="modal modal-lg">
    <div class="mh"><div class="mt" id="tm-title">Script Timing</div><button class="mx" onclick="closeModal('timingModal')">✕</button></div>
    <div class="mb" style="display:grid;gap:10px">
      <div class="frow">
        <div class="fg"><div class="fl">Script</div><div id="tm-script-name" style="font-weight:500"></div></div>
        <div class="fg"><div class="fl">Confidence</div><div id="tm-conf"></div></div>
      </div>
      <div class="fg"><div class="fl">Description</div><div id="tm-desc" style="font-size:11px;color:var(--txd)"></div></div>
      <div id="tm-no-data" style="font-family:var(--mono);font-size:10px;color:var(--txm);padding:8px 0">
        No timing data yet — run this script on a source to collect samples.
      </div>
      <div id="tm-data" style="display:none;display:grid;gap:8px">
        <div style="background:var(--bg3);border:1px solid var(--b1);border-radius:var(--r);padding:8px 10px;display:grid;grid-template-columns:repeat(3,1fr);gap:8px">
          <div class="fg"><div class="fl">Avg Throughput</div><div id="tm-bps" style="font-family:var(--mono);font-size:13px;color:var(--ac2)"></div></div>
          <div class="fg"><div class="fl">Avg Lines/sec</div><div id="tm-lps" style="font-family:var(--mono);font-size:13px;color:var(--ac3)"></div></div>
          <div class="fg"><div class="fl">Samples</div><div id="tm-samples" style="font-family:var(--mono);font-size:13px;color:var(--txd)"></div></div>
        </div>
        <div class="fg">
          <div class="fl">Estimated runtime per source</div>
          <table class="ftable" id="tm-est-table" style="margin-top:4px">
            <thead><tr><th>Source</th><th>Size</th><th>LOC</th><th>Est. Time (bytes)</th><th>Est. Time (lines)</th></tr></thead>
            <tbody id="tm-est-body"></tbody>
          </table>
        </div>
        <div class="fg">
          <div class="fl">Past jobs with this script</div>
          <table class="ftable" id="tm-jobs-table" style="margin-top:4px">
            <thead><tr><th>Source</th><th>Status</th><th>Duration</th><th>Findings</th><th>When</th></tr></thead>
            <tbody id="tm-jobs-body"></tbody>
          </table>
        </div>
      </div>
    </div>
    <div class="mf"><button class="btn btn-s" onclick="closeModal('timingModal')">Close</button></div>
  </div>
</div>

<script>
'use strict';
// ════════════════════════════════════════════════
// State
// ════════════════════════════════════════════════
let scripts=[], sources=[], groups=[], jobs=[];
let currentPage='dashboard';
let editingScriptId=null, editingSourceId=null, editingGroupId=null;
let currentJobId=null, currentSSE=null;
let allFindings=[];
let activeVTab='graph';
let _fmJobId=null, _fmScriptId=null, _fmSourceId=null;
let _jdScriptId=null, _jdSourceId=null;

// ════════════════════════════════════════════════
// Navigation
// ════════════════════════════════════════════════
const PAGE_ORDER=['dashboard','jobs','findings','scripts','sources'];
function showPage(name){
  document.querySelectorAll('.page').forEach(p=>p.classList.remove('active'));
  document.querySelectorAll('.tab').forEach(t=>t.classList.remove('active'));
  const pg=document.getElementById('page-'+name);
  if(pg) pg.classList.add('active');
  const i=PAGE_ORDER.indexOf(name);
  if(i>=0) document.querySelectorAll('.tab')[i].classList.add('active');
  currentPage=name;
  if(name==='dashboard') refreshDashboard();
  else if(name==='jobs') refreshJobs();
  else if(name==='findings') refreshFindings();
  else if(name==='scripts') refreshScripts();
  else if(name==='sources') refreshSources();
}

// ════════════════════════════════════════════════
// API
// ════════════════════════════════════════════════
async function api(method,path,body){
  const o={method,headers:{'Content-Type':'application/json'}};
  if(body) o.body=JSON.stringify(body);
  const r=await fetch(path,o);
  if(r.status===204) return null;
  try{ return await r.json(); }catch(e){ return null; }
}

// ════════════════════════════════════════════════
// Refresh
// ════════════════════════════════════════════════
async function refreshAll(){
  [scripts,sources,groups,jobs]=await Promise.all([
    api('GET','/api/scripts'), api('GET','/api/sources'),
    api('GET','/api/groups'),  api('GET','/api/jobs'),
  ]);
  scripts=scripts||[]; sources=sources||[]; groups=groups||[]; jobs=jobs||[];
  buildAllFindings(); updateStats();
}
async function refreshDashboard(){ await refreshAll(); renderDashJobList(); fetchMetrics(); }
async function refreshJobs(){
  jobs=await api('GET','/api/jobs')||[];
  buildAllFindings(); updateStats(); renderJobList();
}
async function refreshScripts(){
  [scripts,groups]=await Promise.all([api('GET','/api/scripts'),api('GET','/api/groups')]);
  scripts=scripts||[]; groups=groups||[];
  renderScriptTree(); updateStats();
}
async function refreshSources(){ sources=await api('GET','/api/sources')||[]; renderSourceList(); updateStats(); }
async function refreshFindings(){ await refreshAll(); populateFindingsJobSel(); loadFindings(); }

function buildAllFindings(){
  allFindings=[];
  for(const j of jobs){
    if(!j.findings?.length) continue;
    const sc=scripts.find(s=>s.id===j.script_id);
    for(const f of j.findings){
      allFindings.push({...f,
        jobId:j.id, jobName:j.script_name+' → '+j.source_name,
        scriptId:j.script_id, scriptName:j.script_name,
        scriptDesc:sc?.description||'',
        scriptConf:j.script_confidence,
        sourceId:j.source_id, sourceName:j.source_name,
      });
    }
  }
}

function updateStats(){
  const running=jobs.filter(j=>j.status==='running').length;
  const queued=jobs.filter(j=>j.status==='queued').length;
  document.getElementById('stat-jobs').textContent=jobs.length;
  document.getElementById('stat-running').textContent=running;
  document.getElementById('stat-queued').textContent=queued;
  document.getElementById('stat-findings').textContent=allFindings.length;
  document.getElementById('stat-scripts').textContent=scripts.length;
  document.getElementById('stat-sources').textContent=sources.length;
  const tr=document.getElementById('tabRunning');
  if(running>0){tr.textContent=running;tr.style.display='';}else tr.style.display='none';
}

async function fetchMetrics(){
  const m=await api('GET','/api/metrics');
  if(!m) return;
  document.getElementById('m-max').textContent=m.max_concurrent;
  document.getElementById('m-run').textContent=m.running;
  document.getElementById('m-que').textContent=m.queued;
  document.getElementById('m-comp').textContent=m.total_completed;
  document.getElementById('m-find').textContent=m.total_findings;
  const inp=document.getElementById('maxJobsInput');
  if(inp&&!inp.dataset.touched) inp.value=m.max_concurrent;
  document.getElementById('m-eta').textContent=computeQueueETA(m.max_concurrent);
}

// Compute an ETA string for all active (running+queued) jobs.
// Simulates a greedy parallel scheduler. Unknown-duration jobs are assigned
// to slots but their time is tracked separately — output: "~2m30s + ??? for 4 unknown"
function computeQueueETA(maxConcurrent){
  const active=jobs.filter(j=>j.status==='running'||j.status==='pending'||j.status==='queued');
  if(!active.length) return '–';

  const estimates=active.map(j=>{
    const sc=scripts.find(s=>s.id===j.script_id);
    const src=sources.find(s=>s.id===j.source_id);
    const bps=sc?.avg_bytes_per_sec||0;
    const sz=(j.source_size_bytes)||src?.size_bytes||0;
    if(bps>0&&sz>0) return {known:true, secs:sz/bps};
    // Fallback: avg duration of past completed jobs for this script
    const pastDurs=jobs.filter(x=>x.script_id===j.script_id&&x.duration_sec>0).map(x=>x.duration_sec);
    if(pastDurs.length){
      const avg=pastDurs.reduce((a,b)=>a+b,0)/pastDurs.length;
      return {known:true, secs:avg};
    }
    return {known:false, secs:0};
  });

  const n=Math.max(1,maxConcurrent);
  // Two parallel slot arrays: one tracks known time, one tracks unknown count per slot
  const knownSlots=Array(n).fill(0);
  const unknownSlots=Array(n).fill(0);

  for(const e of estimates){
    // Pick the slot that finishes earliest (by known time)
    let best=0;
    for(let i=1;i<n;i++) if(knownSlots[i]<knownSlots[best]) best=i;
    if(e.known) knownSlots[best]+=e.secs;
    else unknownSlots[best]++;
  }

  const totalKnown=Math.max(...knownSlots);
  const totalUnknown=unknownSlots.reduce((a,b)=>a+b,0);

  if(totalKnown===0&&totalUnknown===0) return '–';
  const knownStr=totalKnown>0?'~'+fmtEst(totalKnown):'';
  if(totalUnknown===0) return knownStr;
  if(totalKnown===0) return '??? ('+totalUnknown+' unknown)';
  return knownStr+' + ??? for '+totalUnknown+' unknown';
}

async function setMaxJobs(){
  const v=parseInt(document.getElementById('maxJobsInput').value);
  if(v>0) await api('PUT','/api/config',{max_concurrent:v});
  document.getElementById('maxJobsInput').dataset.touched='';
}

// ════════════════════════════════════════════════
// Helpers
// ════════════════════════════════════════════════
function esc(s){return String(s||'').replace(/&/g,'&amp;').replace(/</g,'&lt;').replace(/>/g,'&gt;').replace(/"/g,'&quot;');}
function badge(st){return '<span class="bdg bdg-'+st+'">'+st+'</span>';}
function ago(ts){
  if(!ts) return '–';
  const d=(Date.now()-new Date(ts))/1000;
  if(d<60) return Math.round(d)+'s ago';
  if(d<3600) return Math.round(d/60)+'m ago';
  if(d<86400) return Math.round(d/3600)+'h ago';
  return Math.round(d/86400)+'d ago';
}
function fmtBytes(b){
  if(!b) return '–';
  if(b>1e9) return (b/1e9).toFixed(1)+' GB';
  if(b>1e6) return (b/1e6).toFixed(1)+' MB';
  if(b>1e3) return (b/1e3).toFixed(1)+' KB';
  return b+' B';
}
function fmtLines(n){if(!n)return '–';if(n>1e6)return (n/1e6).toFixed(1)+'M LOC';if(n>1e3)return (n/1e3).toFixed(0)+'K LOC';return n+' LOC';}
function confPill(v){
  const cls=v>=75?'conf-hi':v>=40?'conf-mid':'conf-lo';
  return '<span class="conf '+cls+'">'+v+'%</span>';
}
function fmtEst(sec){if(sec<60)return Math.round(sec)+'s';if(sec<3600)return (sec/60).toFixed(1)+'m';return (sec/3600).toFixed(1)+'h';}

// Script group hierarchy for a given script ID
function scriptGroupPath(scriptId){
  const path=[];
  function walk(gid,chain){
    const g=groups.find(x=>x.id===gid);
    if(!g) return;
    const newChain=[...chain,g.name];
    if((g.script_ids||[]).includes(scriptId)){path.push(newChain.join(' › '));return;}
    for(const cid of (g.child_ids||[])) walk(cid,newChain);
  }
  const childSet=new Set(groups.flatMap(g=>g.child_ids||[]));
  for(const g of groups) if(!childSet.has(g.id)) walk(g.id,[]);
  return path;
}

// ════════════════════════════════════════════════
// Job items
// ════════════════════════════════════════════════
function jobItemHTML(j){
  const fc=j.findings?.length||0;
  const dur=j.duration_sec?j.duration_sec.toFixed(1)+'s':'–';
  const canRestart=j.status==='failed'||j.status==='canceled';
  const restartBtn=canRestart?'<button class="btn btn-p" style="font-size:9px;padding:2px 5px" onclick="event.stopPropagation();restartJob(\''+j.id+'\')">↺</button>':'';
  const fcBadge=fc?'<span style="font-family:var(--mono);font-size:9px;color:var(--red);white-space:nowrap">'+fc+'✦</span>':'';
  return '<div class="item" style="padding:4px 10px;gap:6px" onclick="openJobDetail(\''+j.id+'\')">'+
    '<div style="font-family:var(--mono);font-size:9px;color:var(--txm);flex-shrink:0;width:52px;text-align:right">'+ago(j.created_at)+'</div>'+
    '<div style="flex:1;min-width:0;font-size:11px;font-weight:500;white-space:nowrap;overflow:hidden;text-overflow:ellipsis">'+esc(j.script_name||'?')+' <span style="color:var(--txm)">→</span> '+esc(j.source_name||'?')+(j.description?' <span style="color:var(--txm);font-weight:400;font-size:10px">'+esc(j.description)+'</span>':'')+'</div>'+
    '<div style="display:flex;gap:4px;align-items:center;flex-shrink:0">'+
      '<span style="font-family:var(--mono);font-size:9px;color:var(--txd)">'+dur+'</span>'+
      fcBadge+
      restartBtn+
      confPill(j.script_confidence||0)+
      badge(j.status)+
    '</div>'+
  '</div>';
}
function renderDashJobList(){
  const el=document.getElementById('dashJobList');
  el.innerHTML=jobs.length?jobs.slice(0,14).map(jobItemHTML).join(''):'<div class="empty"><div class="empty-ico">◈</div>No jobs</div>';
}
function renderJobList(){
  const el=document.getElementById('jobList');
  el.innerHTML=jobs.length?jobs.map(jobItemHTML).join(''):'<div class="empty"><div class="empty-ico">◈</div>No jobs</div>';
}

// ════════════════════════════════════════════════
// Script Tree (nested details/summary)
// ════════════════════════════════════════════════
function getScriptFilter(){
  const q=(document.getElementById('scriptSearch').value||'').toLowerCase();
  const lang=document.querySelector('[name="lf"]:checked')?.value||'';
  const minConf=parseInt(document.getElementById('confFilter').value)||0;
  return {q,lang,minConf};
}

function scriptMatchesFilter(sc,f){
  if(f.lang && sc.language!==f.lang) return false;
  if(sc.confidence<f.minConf) return false;
  if(f.q && !sc.name.toLowerCase().includes(f.q) && !(sc.description||'').toLowerCase().includes(f.q)) return false;
  return true;
}

// Count all resolved scripts under a group (recursively), respecting filters
function countGroupScripts(gid, f){
  const seen=new Set();
  function walk(gid){
    if(seen.has('g:'+gid)) return 0;
    seen.add('g:'+gid);
    const g=groups.find(x=>x.id===gid);
    if(!g) return 0;
    let n=0;
    for(const sid of (g.script_ids||[])){
      if(seen.has('s:'+sid)) continue;
      seen.add('s:'+sid);
      const sc=scripts.find(s=>s.id===sid);
      if(sc && scriptMatchesFilter(sc,f)) n++;
    }
    for(const cid of (g.child_ids||[])) n+=walk(cid);
    return n;
  }
  return walk(gid);
}

function renderScriptRow(sc){
  const fc=allFindings.filter(f=>f.scriptId===sc.id).length;
  const fcBadge=fc?'<span style="font-family:var(--mono);font-size:9px;color:var(--red);white-space:nowrap" title="'+fc+' findings total">'+fc+'✦</span>':'';
  return '<div class="sc-row" id="scrow-'+sc.id+'">'+
    '<div class="iico" style="background:rgba(32,200,224,.07);color:var(--ac2)">{ }</div>'+
    '<div class="sc-row-body">'+
      '<div class="iname">'+esc(sc.name)+'</div>'+
      '<div class="imeta">'+esc(sc.description||'—')+'</div>'+
    '</div>'+
    '<div class="iact">'+
      '<span class="bdg bdg-'+sc.language+'">'+sc.language.toUpperCase()+'</span>'+
      confPill(sc.confidence||0)+
      fcBadge+
      (sc.avg_bytes_per_sec?'<span class="est-time" title="avg throughput">'+fmtBytes(sc.avg_bytes_per_sec)+'/s</span>':'')+
      '<button class="btn btn-a" style="font-size:9px;padding:2px 5px" title="Run on sources…" onclick="event.stopPropagation();openScriptRunModal(\''+sc.id+'\')">▶ …</button>'+
      '<button class="btn btn-s" style="font-size:9px" onclick="event.stopPropagation();openTimingModal(\''+sc.id+'\')">⏱</button>'+
      '<button class="btn btn-s" style="font-size:9px" onclick="event.stopPropagation();openScriptModal(\''+sc.id+'\')">Edit</button>'+
      '<button class="btn btn-s" style="font-size:9px" onclick="event.stopPropagation();dupScript(\''+sc.id+'\')">Dup</button>'+
      '<button class="btn btn-d" style="font-size:9px" onclick="event.stopPropagation();deleteScript(\''+sc.id+'\')">Del</button>'+
    '</div>'+
  '</div>';
}

function renderGroupDetails(gid, f, depth=0){
  const g=groups.find(x=>x.id===gid);
  if(!g) return '';
  const cnt=countGroupScripts(gid,f);
  if(cnt===0 && f.q) return ''; // hide empty groups when filtering

  let inner='';
  // child groups
  for(const cid of (g.child_ids||[])){
    inner+=renderGroupDetails(cid,f,depth+1);
  }
  // scripts in this group
  for(const sid of (g.script_ids||[])){
    const sc=scripts.find(s=>s.id===sid);
    if(sc && scriptMatchesFilter(sc,f)) inner+=renderScriptRow(sc);
  }
  if(!inner) return '';

  const openAttr=f.q?'open':'';
  return '<details class="gd" '+openAttr+' id="gd-'+g.id+'">'+
    '<summary>'+
      '<span class="sum-arr">▸</span>'+
      '<span class="iico" style="color:var(--ac4);background:rgba(176,96,248,.07)">▦</span>'+
      '<span class="sum-lbl">'+esc(g.name)+(g.description?' <span style="color:var(--txm);font-weight:400;font-size:10px">'+esc(g.description)+'</span>':'')+'</span>'+
      '<span class="sum-cnt">'+cnt+' script'+(cnt!==1?'s':'')+'</span>'+
      '<div style="display:flex;gap:3px;flex-shrink:0" onclick="event.stopPropagation()">'+
        '<button class="btn btn-a" style="font-size:9px;padding:2px 5px" title="Run on all sources" onclick="runGroupOnAllSources(\''+g.id+'\')">▶ All</button>'+
        '<button class="btn btn-2" style="font-size:9px;padding:2px 5px" title="Choose sources" onclick="openGroupRunModalForGroup(\''+g.id+'\')">▶ …</button>'+
        '<button class="btn btn-s" style="font-size:9px;padding:2px 5px" onclick="openGroupModal(\''+g.id+'\')">Edit</button>'+
        '<button class="btn btn-d" style="font-size:9px;padding:2px 5px" onclick="deleteGroup(\''+g.id+'\')">Del</button>'+
      '</div>'+
    '</summary>'+
    '<div class="gd-body">'+inner+'</div>'+
  '</details>';
}

function renderScriptTree(){
  const el=document.getElementById('scriptTree');
  const f=getScriptFilter();
  const sortMode=document.getElementById('scriptSort').value;

  // When a non-tree sort is chosen, render all matching scripts as a flat sorted list
  if(sortMode!=='tree'){
    let list=scripts.filter(sc=>scriptMatchesFilter(sc,f));
    const findingsForScript=id=>allFindings.filter(x=>x.scriptId===id).length;
    if(sortMode==='findings') list.sort((a,b)=>findingsForScript(b.id)-findingsForScript(a.id));
    else if(sortMode==='fastest') list.sort((a,b)=>(b.avg_bytes_per_sec||0)-(a.avg_bytes_per_sec||0));
    else if(sortMode==='slowest') list.sort((a,b)=>(a.avg_bytes_per_sec||Infinity)-(b.avg_bytes_per_sec||Infinity));
    else if(sortMode==='conf') list.sort((a,b)=>(b.confidence||0)-(a.confidence||0));
    else if(sortMode==='name') list.sort((a,b)=>a.name.localeCompare(b.name));
    el.innerHTML=list.length?list.map(renderScriptRow).join(''):'<div class="empty"><div class="empty-ico">◈</div>No scripts match</div>';
    return;
  }

  // Tree order
  const childSet=new Set(groups.flatMap(g=>g.child_ids||[]));
  const topGroups=groups.filter(g=>!childSet.has(g.id));
  const groupedScriptIds=new Set(groups.flatMap(g=>g.script_ids||[]));
  let html='';
  for(const g of topGroups) html+=renderGroupDetails(g.id,f,0);
  const ungrouped=scripts.filter(sc=>!groupedScriptIds.has(sc.id)&&scriptMatchesFilter(sc,f));
  if(ungrouped.length){
    html+='<div class="ungrouped-hdr">Ungrouped ('+ungrouped.length+')</div>';
    html+=ungrouped.map(renderScriptRow).join('');
  }
  el.innerHTML=html||'<div class="empty"><div class="empty-ico">◈</div>No scripts match</div>';
}

document.querySelectorAll('[name="lf"]').forEach(r=>r.addEventListener('change',()=>{
  document.querySelectorAll('[name="lf"]').forEach((x,i)=>{
    x.parentElement.className='ropt'+(x.checked?' sel':'');
  });
  renderScriptTree();
}));

// ════════════════════════════════════════════════
// Source List
// ════════════════════════════════════════════════
function renderSourceList(){
  const el=document.getElementById('sourceList');
  if(!sources.length){el.innerHTML='<div class="empty"><div class="empty-ico">◈</div>No sources</div>';return;}
  el.innerHTML='<table class="ftable" style="table-layout:fixed">'+
    '<thead><tr>'+
      '<th style="width:160px">Name</th>'+
      '<th>Path</th>'+
      '<th style="width:90px">Size</th>'+
      '<th style="width:70px">Lines</th>'+
      '<th style="width:60px">Tags</th>'+
      '<th style="width:170px"></th>'+
    '</tr></thead>'+
    '<tbody>'+sources.map(s=>{
      const sz=s.size_bytes?fmtBytes(s.size_bytes):'–';
      const lc=s.line_count?fmtLines(s.line_count):'–';
      const tags=(s.tags||[]).map(t=>'<span class="bdg bdg-'+t+'">'+t.toUpperCase()+'</span>').join('');
      const pathParts=s.path.replace(/\\/g,'/').split('/');
      const shortPath=pathParts.length>4?'…/'+pathParts.slice(-3).join('/'):s.path;
      return '<tr>'+
        '<td style="font-weight:500;font-size:11px;white-space:nowrap;overflow:hidden;text-overflow:ellipsis;max-width:160px" title="'+esc(s.name)+'">'+esc(s.name)+'</td>'+
        '<td style="font-family:var(--mono);font-size:9px;color:var(--txd);white-space:nowrap;overflow:hidden;text-overflow:ellipsis" title="'+esc(s.path)+'">'+esc(shortPath)+'</td>'+
        '<td style="font-family:var(--mono);font-size:10px;color:var(--txm)">'+sz+'</td>'+
        '<td style="font-family:var(--mono);font-size:10px;color:var(--txm)">'+lc+'</td>'+
        '<td>'+tags+'</td>'+
        '<td style="text-align:right"><div style="display:flex;gap:3px;justify-content:flex-end">'+
          '<button class="btn btn-2" style="font-size:9px;padding:2px 5px" onclick="sourceRunAll(\''+s.id+'\')">▶ All</button>'+
          '<button class="btn btn-s" style="font-size:9px;padding:2px 5px" onclick="openSourceModal(\''+s.id+'\')">Edit</button>'+
          '<button class="btn btn-s" style="font-size:9px;padding:2px 5px" onclick="dupSource(\''+s.id+'\')">Dup</button>'+
          '<button class="btn btn-d" style="font-size:9px;padding:2px 5px" onclick="deleteSource(\''+s.id+'\')">Del</button>'+
        '</div></td>'+
      '</tr>';
    }).join('')+
    '</tbody></table>';
}

// ════════════════════════════════════════════════
// Script CRUD
// ════════════════════════════════════════════════
function openScriptModal(id){
  editingScriptId=id||null;
  document.getElementById('scriptModalTitle').textContent=id?'Edit Script':'New Script';
  const sc=id?scripts.find(s=>s.id===id):null;
  document.getElementById('sc-name').value=sc?.name||'';
  document.getElementById('sc-desc').value=sc?.description||'';
  document.getElementById('sc-pattern').value=sc?.pattern||'';
  document.getElementById('sc-flags').value=sc?.extra_flags||'';
  const conf=sc?.confidence??50;
  document.getElementById('sc-conf').value=conf;
  document.getElementById('sc-conf-val').textContent=conf;
  // timing info
  const ti=document.getElementById('sc-timing-info');
  if(sc?.timing_samples>0){
    ti.style.display='';
    ti.textContent='Avg: '+fmtBytes(sc.avg_bytes_per_sec)+'/s  ·  '+fmtLines(sc.avg_lines_per_sec)+'/s  ·  '+sc.timing_samples+' sample'+(sc.timing_samples!==1?'s':'');
  } else ti.style.display='none';
  const lang=sc?.language||'c';
  document.querySelectorAll('[name="sc-lang"]').forEach(r=>{r.checked=r.value===lang;r.parentElement.className='ropt'+(r.value===lang?' sel':'');});
  openModal('scriptModal');
}

async function saveScript(){
  const lang=document.querySelector('[name="sc-lang"]:checked').value;
  const sc={name:document.getElementById('sc-name').value.trim(),description:document.getElementById('sc-desc').value.trim(),pattern:document.getElementById('sc-pattern').value,language:lang,extra_flags:document.getElementById('sc-flags').value.trim(),confidence:parseInt(document.getElementById('sc-conf').value)||0};
  if(!sc.name||!sc.pattern){alert('Name and pattern required.');return;}
  if(editingScriptId){sc.id=editingScriptId;await api('PUT','/api/scripts/'+editingScriptId,sc);}
  else await api('POST','/api/scripts',sc);
  closeModal('scriptModal'); await refreshScripts();
}

async function dupScript(id){
  await api('POST','/api/scripts/'+id+'/duplicate');
  await refreshScripts();
}

async function deleteScript(id){if(!confirm('Delete script?'))return;await api('DELETE','/api/scripts/'+id);await refreshScripts();}

// ════════════════════════════════════════════════
// Group CRUD
// ════════════════════════════════════════════════
function openGroupModal(id){
  editingGroupId=id||null;
  document.getElementById('groupModalTitle').textContent=id?'Edit Group':'New Group';
  const g=id?groups.find(x=>x.id===id):null;
  document.getElementById('grp-name').value=g?.name||'';
  document.getElementById('grp-desc').value=g?.description||'';
  const selSc=new Set(g?.script_ids||[]);
  document.getElementById('grp-scripts-list').innerHTML=scripts.length
    ?scripts.map(s=>'<label class="ci"><input type="checkbox" value="'+s.id+'"'+(selSc.has(s.id)?' checked':'')+'>'+esc(s.name)+'<span class="ci-sub"> '+s.language+'</span></label>').join('')
    :'<span style="color:var(--txm);font-family:var(--mono);font-size:10px">No scripts</span>';
  const selCh=new Set(g?.child_ids||[]);
  const avail=groups.filter(x=>x.id!==id);
  document.getElementById('grp-children-list').innerHTML=avail.length
    ?avail.map(x=>'<label class="ci"><input type="checkbox" value="'+x.id+'"'+(selCh.has(x.id)?' checked':'')+'>'+esc(x.name)+'</label>').join('')
    :'<span style="color:var(--txm);font-family:var(--mono);font-size:10px">No other groups</span>';
  openModal('groupModal');
}

async function saveGroup(){
  const name=document.getElementById('grp-name').value.trim();
  if(!name){alert('Name required.');return;}
  const scriptIds=[...document.querySelectorAll('#grp-scripts-list input:checked')].map(i=>i.value);
  const childIds=[...document.querySelectorAll('#grp-children-list input:checked')].map(i=>i.value);
  const g={name,description:document.getElementById('grp-desc').value.trim(),script_ids:scriptIds,child_ids:childIds};
  if(editingGroupId){g.id=editingGroupId;await api('PUT','/api/groups/'+editingGroupId,g);}
  else await api('POST','/api/groups',g);
  closeModal('groupModal'); await refreshScripts();
}

async function deleteGroup(id){if(!confirm('Delete group?'))return;await api('DELETE','/api/groups/'+id);await refreshScripts();}

// ════════════════════════════════════════════════
// Source CRUD
// ════════════════════════════════════════════════
function openSourceModal(id){
  editingSourceId=id||null;
  document.getElementById('sourceModalTitle').textContent=id?'Edit Source':'New Source';
  const s=id?sources.find(x=>x.id===id):null;
  document.getElementById('src-name').value=s?.name||'';
  document.getElementById('src-desc').value=s?.description||'';
  document.getElementById('src-path').value=s?.path||'';
  const tags=new Set(s?.tags||[]);
  document.querySelectorAll('[name="src-tag"]').forEach(cb=>{cb.checked=tags.has(cb.value);cb.parentElement.className='ropt'+(tags.has(cb.value)?' sel':'');});
  const sd=document.getElementById('src-stats-display');
  if(s?.size_bytes){sd.style.display='';sd.textContent=fmtBytes(s.size_bytes)+' · '+fmtLines(s.line_count)+(s.last_scanned?' · scanned '+ago(s.last_scanned):'');}
  else sd.style.display='none';
  openModal('sourceModal');
}

document.querySelectorAll('[name="src-tag"]').forEach(cb=>cb.addEventListener('change',()=>{
  document.querySelectorAll('[name="src-tag"]').forEach(x=>{x.parentElement.className='ropt'+(x.checked?' sel':'');});
}));

async function saveSource(){
  const tags=[...document.querySelectorAll('[name="src-tag"]:checked')].map(c=>c.value);
  const s={name:document.getElementById('src-name').value.trim(),description:document.getElementById('src-desc').value.trim(),path:document.getElementById('src-path').value.trim(),tags};
  if(!s.name||!s.path){alert('Name and path required.');return;}
  if(editingSourceId){s.id=editingSourceId;await api('PUT','/api/sources/'+editingSourceId,s);}
  else await api('POST','/api/sources',s);
  closeModal('sourceModal'); await refreshSources();
}

async function dupSource(id){await api('POST','/api/sources/'+id+'/duplicate');await refreshSources();}
async function deleteSource(id){if(!confirm('Delete source?'))return;await api('DELETE','/api/sources/'+id);await refreshSources();}
async function sourceRunAll(id){
  const r=await api('POST','/api/sources/'+id+'/run-all');
  if(r) { showPage('jobs'); await refreshJobs(); }
}

// ════════════════════════════════════════════════
// Job CRUD — single
// ════════════════════════════════════════════════
function openNewJobModal(){
  document.getElementById('job-script-sel').innerHTML='<option value="">— select script —</option>'+scripts.map(s=>'<option value="'+s.id+'">'+esc(s.name)+'</option>').join('');
  document.getElementById('job-source-sel').innerHTML='<option value="">— select source —</option>'+sources.map(s=>'<option value="'+s.id+'">'+esc(s.name)+' ('+esc(s.path)+')</option>').join('');
  document.getElementById('job-desc').value='';
  document.getElementById('job-cmd-preview').style.display='none';
  document.getElementById('job-est').style.display='none';
  openModal('jobModal');
}

function updateJobPreview(){
  const sid=document.getElementById('job-script-sel').value;
  const srcid=document.getElementById('job-source-sel').value;
  const p=document.getElementById('job-cmd-preview');
  const est=document.getElementById('job-est');
  if(!sid||!srcid){p.style.display='none';est.style.display='none';return;}
  const sc=scripts.find(s=>s.id===sid), src=sources.find(s=>s.id===srcid);
  if(!sc||!src){p.style.display='none';return;}
  p.style.display='block'; p.innerHTML=buildCmdHTML(sc,src);
  // Estimate time
  if(sc.avg_bytes_per_sec>0&&src.size_bytes>0){
    const secs=src.size_bytes/sc.avg_bytes_per_sec;
    est.style.display='';
    est.textContent='~'+fmtEst(secs)+' estimated (based on '+sc.timing_samples+' past run'+(sc.timing_samples!==1?'s':'')+')';
  } else est.style.display='none';
}

function buildCmdHTML(sc,src){
  let f=''; if(sc.language==='cpp'||sc.language==='both') f+=' <span class="cf">--cpp</span>';
  if(sc.extra_flags) f+=' <span class="cf">'+esc(sc.extra_flags)+'</span>';
  return '<span class="cb">weggli</span>'+f+' <span class="cp">\''+esc((sc.pattern||'').replace(/\n/g,'↵').slice(0,80))+'\'</span> <span class="cx">'+esc(src.path)+'</span>';
}

async function startJob(){
  const sid=document.getElementById('job-script-sel').value, srcid=document.getElementById('job-source-sel').value;
  if(!sid||!srcid){alert('Select script and source.');return;}
  document.getElementById('startJobBtn').disabled=true;
  const job=await api('POST','/api/jobs',{script_id:sid,source_id:srcid,description:document.getElementById('job-desc').value.trim()});
  document.getElementById('startJobBtn').disabled=false;
  closeModal('jobModal');
  if(job?.id){jobs.unshift(job);openJobDetail(job.id);}
}

// ════════════════════════════════════════════════
// Group Run
// ════════════════════════════════════════════════
function openGroupRunModal(preGid){
  document.getElementById('grr-group-sel').innerHTML='<option value="">— select group —</option>'+groups.map(g=>'<option value="'+g.id+'">'+esc(g.name)+'</option>').join('');
  const sl=document.getElementById('grr-sources-list');
  sl.innerHTML=sources.map(s=>'<label class="ci"><input type="checkbox" value="'+s.id+'">'+esc(s.name)+' <span class="ci-sub">'+esc(s.path)+'</span></label>').join('');
  if(preGid) document.getElementById('grr-group-sel').value=preGid;
  document.getElementById('grr-preview').style.display='none';
  updateGroupRunPreview();
  openModal('groupRunModal');
}
function openGroupRunModalForGroup(gid){ openGroupRunModal(gid); }

function resolveGroupScripts(gid,seen=new Set()){
  if(seen.has('g:'+gid)) return [];
  seen.add('g:'+gid);
  const g=groups.find(x=>x.id===gid); if(!g) return [];
  const r=[];
  for(const sid of (g.script_ids||[])) if(!seen.has('s:'+sid)){seen.add('s:'+sid);r.push(sid);}
  for(const cid of (g.child_ids||[])) r.push(...resolveGroupScripts(cid,seen));
  return r;
}

function updateGroupRunPreview(){
  const gid=document.getElementById('grr-group-sel').value;
  const p=document.getElementById('grr-preview');
  if(!gid){p.style.display='none';return;}
  const g=groups.find(x=>x.id===gid); if(!g){p.style.display='none';return;}
  const sids=resolveGroupScripts(gid);
  const selectedSrcIds=[...document.querySelectorAll('#grr-sources-list input:checked')].map(i=>i.value);
  p.style.display='block';
  p.innerHTML='<span style="color:var(--ac4)">'+sids.length+'</span> scripts × <span style="color:var(--ac2)">'+(selectedSrcIds.length||'?')+'</span> sources = <span style="color:var(--ac)">'+sids.length*(selectedSrcIds.length||0)+'</span> jobs<br>'+
    sids.map(sid=>{const s=scripts.find(x=>x.id===sid);return s?'<span style="color:var(--ac)">'+esc(s.name)+'</span>':'?';}).join(', ');
}

document.getElementById('grr-group-sel').addEventListener('change',updateGroupRunPreview);
document.getElementById('grr-sources-list').addEventListener('change',updateGroupRunPreview);

async function runGroup(){
  const gid=document.getElementById('grr-group-sel').value;
  const srcIds=[...document.querySelectorAll('#grr-sources-list input:checked')].map(i=>i.value);
  if(!gid||!srcIds.length){alert('Select group and at least one source.');return;}
  document.getElementById('runGroupBtn').disabled=true;
  const res=await api('POST','/api/groups/'+gid+'/run',{source_ids:srcIds});
  document.getElementById('runGroupBtn').disabled=false;
  closeModal('groupRunModal');
  if(res?.jobs?.length){await refreshJobs();showPage('jobs');}
}

function openScriptRunModal(preScriptId){
  const sel=document.getElementById('srr-script-sel');
  sel.innerHTML='<option value="">— select script —</option>'+scripts.map(s=>'<option value="'+s.id+'">'+esc(s.name)+'</option>').join('');
  if(preScriptId) sel.value=preScriptId;
  const sl=document.getElementById('srr-sources-list');
  sl.innerHTML=sources.map(s=>'<label class="ci"><input type="checkbox" value="'+s.id+'">'+esc(s.name)+' <span class="ci-sub">'+esc(s.path)+'</span></label>').join('');
  document.getElementById('srr-preview').style.display='none';
  updateScriptRunPreview();
  openModal('scriptRunModal');
}
function updateScriptRunPreview(){
  const sid=document.getElementById('srr-script-sel').value;
  const p=document.getElementById('srr-preview');
  if(!sid){p.style.display='none';return;}
  const sc=scripts.find(s=>s.id===sid);
  const selectedSrcIds=[...document.querySelectorAll('#srr-sources-list input:checked')].map(i=>i.value);
  p.style.display='block';
  p.innerHTML='<span style="color:var(--ac2)">'+esc(sc?.name||'?')+'</span> on <span style="color:var(--ac4)">'+(selectedSrcIds.length||'?')+'</span> source'+(selectedSrcIds.length!==1?'s':'')
    +(sc?'<br><span style="color:var(--txm)">'+esc(sc.pattern)+'</span>':'');
}
document.getElementById('srr-script-sel').addEventListener('change',updateScriptRunPreview);
document.getElementById('srr-sources-list').addEventListener('change',updateScriptRunPreview);
async function runScriptOnSources(){
  const sid=document.getElementById('srr-script-sel').value;
  const srcIds=[...document.querySelectorAll('#srr-sources-list input:checked')].map(i=>i.value);
  if(!sid||!srcIds.length){alert('Select a script and at least one source.');return;}
  document.getElementById('runScriptBtn').disabled=true;
  const sc=scripts.find(s=>s.id===sid);
  for(const srcId of srcIds){
    const src=sources.find(s=>s.id===srcId);
    await api('POST','/api/jobs',{script_id:sid,source_id:srcId,description:(sc?.name||'')+' → '+(src?.name||'')});
  }
  document.getElementById('runScriptBtn').disabled=false;
  closeModal('scriptRunModal');
  await refreshJobs(); showPage('jobs');
}

async function runGroupOnAllSources(gid){
  if(!sources.length){alert('No sources defined.');return;}
  const g=groups.find(x=>x.id===gid);
  const sids=resolveGroupScripts(gid);
  if(!sids.length){alert('Group has no scripts.');return;}
  const n=sids.length*sources.length;
  if(!confirm('Run '+sids.length+' script'+(sids.length!==1?'s':'')+' × '+sources.length+' source'+(sources.length!==1?'s':'')+' = '+n+' job'+(n!==1?'s':'')+' total?')) return;
  const srcIds=sources.map(s=>s.id);
  const res=await api('POST','/api/groups/'+gid+'/run',{source_ids:srcIds});
  if(res?.jobs?.length){await refreshJobs();showPage('jobs');}
}

// ════════════════════════════════════════════════
// Job Detail
// ════════════════════════════════════════════════
async function openJobDetail(id){
  currentJobId=id;
  if(currentSSE){currentSSE.close();currentSSE=null;}
  document.querySelectorAll('.page').forEach(p=>p.classList.remove('active'));
  document.querySelectorAll('.tab').forEach(t=>t.classList.remove('active'));
  document.getElementById('page-job-detail').classList.add('active');
  currentPage='job-detail';
  const job=await api('GET','/api/jobs/'+id);
  if(!job) return;
  renderJobDetail(job);
  if(job.status==='running'||job.status==='pending'){startStream(id);document.getElementById('jd-cancel-btn').style.display='';}
  else document.getElementById('jd-cancel-btn').style.display='none';
  const canRestart=job.status==='failed'||job.status==='canceled';
  document.getElementById('jd-restart-btn').style.display=canRestart?'':'none';
  document.getElementById('jd-findings-btn').style.display=(job.findings?.length>0)?'':'none';
}

function renderJobDetail(job){
  const sc=scripts.find(s=>s.id===job.script_id);
  const src=sources.find(s=>s.id===job.source_id);
  _jdScriptId=job.script_id; _jdSourceId=job.source_id;
  document.getElementById('jd-title').textContent=(job.script_name||job.script_id)+' → '+(job.source_name||job.source_id);
  document.getElementById('jd-badge').innerHTML=badge(job.status);
  document.getElementById('jd-script-lnk').textContent=job.script_name||(sc?.name||job.script_id);
  document.getElementById('jd-script-desc').textContent=sc?.description||'';
  const metaEl=document.getElementById('jd-script-meta');
  if(sc) metaEl.innerHTML='<span class="bdg bdg-'+sc.language+'">'+sc.language.toUpperCase()+'</span>'+confPill(sc.confidence||0);
  else metaEl.innerHTML=confPill(job.script_confidence||0);
  document.getElementById('jd-script-pattern').textContent=sc?.pattern||'';
  document.getElementById('jd-source-lnk').textContent=(job.source_name||(src?.name||job.source_id))+' ('+(job.source_path||(src?.path||''))+')';
  document.getElementById('jd-started').textContent=job.started_at?new Date(job.started_at).toLocaleString():'–';
  let dur='–';
  if(job.duration_sec) dur=job.duration_sec.toFixed(1)+'s';
  else if(job.status==='running') dur='running…';
  document.getElementById('jd-duration').textContent=dur;
  const fc=job.findings?.length||0;
  document.getElementById('jd-fcount').textContent=fc>0?fc+' findings':'–';
  document.getElementById('jd-desc').textContent=job.description||'–';
  const szBytes=job.source_size_bytes||src?.size_bytes||0;
  const szLines=job.source_line_count||src?.line_count||0;
  document.getElementById('jd-srcsize').textContent=szBytes?fmtBytes(szBytes)+' · '+fmtLines(szLines):'–';
  if(sc&&src) document.getElementById('jd-cmd').innerHTML=buildCmdHTML(sc,src);
  else document.getElementById('jd-cmd').textContent='weggli …';
  const out=document.getElementById('jd-out');
  out.innerHTML='';
  if(job.output) appendOutputLines(out, job.output);
  out.scrollTop=out.scrollHeight;
  const spin=document.getElementById('jd-spin'), tt=document.getElementById('jd-ttitle');
  if(job.status==='running'||job.status==='pending'||job.status==='queued'){
    spin.style.display=''; tt.textContent=job.status==='queued'?'queued…':'streaming…';
  } else {
    spin.style.display='none'; tt.textContent='output ('+job.status+', exit '+job.exit_code+', '+dur+')';
  }
}

function jdNavScript(e){
  e.preventDefault();
  if(_jdScriptId){ closeModal('scriptModal'); openScriptModal(_jdScriptId); }
  return false;
}
function jdNavSource(e){
  e.preventDefault();
  if(_jdSourceId){ openSourceModal(_jdSourceId); }
  return false;
}

// JS-side ANSI SGR → HTML converter. Mirrors ansi.go so old stored output
// with raw escape codes still renders correctly, and new output (already HTML
// from the Go side) passes through unchanged (no ANSI codes to find).
const SGR_COLORS={
  30:'#4a5568',31:'#fc8181',32:'#68d391',33:'#f6e05e',
  34:'#63b3ed',35:'#b794f4',36:'#76e4f7',37:'#e2e8f0',
  90:'#718096',91:'#fc8181',92:'#68d391',93:'#faf089',
  94:'#76e4f7',95:'#fbb6ce',96:'#81e6d9',97:'#f7fafc'
};
function ansiLineToHTML(line){
  // Detect already-processed HTML (from Go's ansiToHTML): contains entities or tags.
  // In that case pass through as-is — it's already safe HTML.
  if(line.includes('<span') || line.includes('&lt;') || line.includes('&amp;') || line.includes('&gt;')) return line;
  // Plain text or raw ANSI: HTML-escape first, then convert ANSI SGR codes.
  let s=line.replace(/&/g,'&amp;').replace(/</g,'&lt;').replace(/>/g,'&gt;');
  let open=0;
  s=s.replace(/\x1b\[([0-9;]*)m/g,(_,p)=>{
    if(!p||p==='0'){
      const close='</span>'.repeat(open); open=0; return close;
    }
    let bold=false,color='',bg='';
    for(const c of p.split(';')){
      const n=parseInt(c);
      if(n===1)bold=true;
      else if(SGR_COLORS[n])color=SGR_COLORS[n];
    }
    if(!bold&&!color&&!bg) return '';
    let style='';
    if(bold)style+='font-weight:700;';
    if(color)style+='color:'+color+';';
    open++;
    return '<span style="'+style+'">';
  });
  // Drop any other non-SGR escape sequences
  s=s.replace(/\x1b\[[0-9;]*[ABCDHKJST]/g,'');
  return s+'</span>'.repeat(open);
}

// appendOutputLines splits HTML/ANSI output at newlines and appends each line
// as an individual <div class="tline"> — safe for streaming since each line
// is an independent HTML fragment with no cross-line open tags.
function appendOutputLines(container, raw){
  const lines=raw.split('\n');
  const frag=document.createDocumentFragment();
  for(let i=0;i<lines.length;i++){
    const line=lines[i];
    if(i===lines.length-1 && line==='') continue; // trailing newline artifact
    const div=document.createElement('div');
    div.className='tline';
    div.innerHTML=ansiLineToHTML(line);
    frag.appendChild(div);
  }
  container.appendChild(frag);
}

function startStream(id){
  const out=document.getElementById('jd-out');
  const sse=new EventSource('/api/jobs/stream/'+id);
  currentSSE=sse;
  sse.onmessage=e=>{
    // Each SSE data payload is one line of HTML (newlines encoded as \n literal).
    // Decode \n back, split into lines, append each as a separate .tline div.
    const text=e.data.replace(/\\n/g,'\n');
    appendOutputLines(out, text);
    out.scrollTop=out.scrollHeight;
  };
  sse.addEventListener('done',async()=>{
    sse.close(); currentSSE=null;
    document.getElementById('jd-spin').style.display='none';
    document.getElementById('jd-cancel-btn').style.display='none';
    const job=await api('GET','/api/jobs/'+id);
    if(job){
      document.getElementById('jd-badge').innerHTML=badge(job.status);
      const dur=job.duration_sec?job.duration_sec.toFixed(1)+'s':'–';
      document.getElementById('jd-ttitle').textContent='output ('+job.status+', exit '+job.exit_code+', '+dur+')';
      document.getElementById('jd-duration').textContent=dur;
      const fc=job.findings?.length||0;
      document.getElementById('jd-fcount').textContent=fc>0?fc+' findings':'–';
      document.getElementById('jd-findings-btn').style.display=fc>0?'':'none';
    }
    jobs=await api('GET','/api/jobs')||[]; buildAllFindings(); updateStats();
  });
  sse.onerror=()=>{sse.close();currentSSE=null;};
}

async function cancelCurrentJob(){ if(currentJobId) await api('POST','/api/jobs/'+currentJobId+'/cancel'); }

async function restartJob(jobId){
  const old=jobs.find(j=>j.id===jobId);
  if(!old) return;
  const job=await api('POST','/api/jobs',{script_id:old.script_id,source_id:old.source_id,description:old.description});
  if(job?.id){ jobs.unshift(job); renderJobList(); openJobDetail(job.id); }
}
async function restartCurrentJob(){ if(currentJobId) await restartJob(currentJobId); }
function openJobFindings(){
  document.getElementById('findings-job-sel').value=currentJobId;
  showPage('findings');
}

// ════════════════════════════════════════════════
// Findings page
// ════════════════════════════════════════════════
function populateFindingsJobSel(){
  const sel=document.getElementById('findings-job-sel'), cur=sel.value;
  const withFindings=jobs.filter(j=>j.findings?.length);
  sel.innerHTML='<option value="">— all jobs —</option>'+
    withFindings.map(j=>'<option value="'+j.id+'">'+esc(j.script_name+' → '+j.source_name)+' ('+j.findings.length+')</option>').join('');
  // Restore previous selection if still valid, otherwise default to first job with findings
  if(cur && withFindings.find(j=>j.id===cur)){
    sel.value=cur;
  } else if(withFindings.length){
    sel.value=withFindings[0].id;
  }
}

function getFilteredFindings(){
  const jid=document.getElementById('findings-job-sel').value;
  return jid?allFindings.filter(f=>f.jobId===jid):allFindings;
}

function loadFindings(){renderFindingTable();renderGraph(getFilteredFindings());}

function switchVTab(name){
  activeVTab=name;
  document.querySelectorAll('.vtab2').forEach((t,i)=>t.classList.toggle('active',['graph','table'][i]===name));
  document.getElementById('vp-graph').style.display=name==='graph'?'flex':'none';
  document.getElementById('vp-table').style.display=name==='table'?'flex':'none';
  if(name==='graph') renderGraph(getFilteredFindings());
}

function renderFindingTable(){
  const findings=getFilteredFindings();
  const q=(document.getElementById('ftFilter').value||'').toLowerCase();
  const filtered=q?findings.filter(f=>
    f.file.toLowerCase().includes(q)||String(f.line).includes(q)||
    (f.snippet||'').toLowerCase().includes(q)||
    (f.scriptName||'').toLowerCase().includes(q)||(f.sourceName||'').toLowerCase().includes(q)
  ):findings;
  document.getElementById('ftCount').textContent=filtered.length+'/'+findings.length;
  const tbody=document.getElementById('ftBody');
  if(!filtered.length){tbody.innerHTML='<tr><td colspan="8" style="text-align:center;padding:20px;color:var(--txm);font-family:var(--mono);font-size:10px">No findings</td></tr>';return;}
  tbody.innerHTML=filtered.map((f,i)=>{
    const parts=f.file.replace(/\\/g,'/').split('/');
    const sf=parts.length>2?'…/'+parts.slice(-2).join('/'):f.file;
    const snipHTML=f.snippet_html
      ? '<span class="fs2">'+f.snippet_html.split('\n')[0].trim().slice(0,300)+'</span>'
      : '<span class="fs2">'+esc((f.snippet||'').split('\n')[0].trim().slice(0,60))+'</span>';
    const idx=allFindings.indexOf(f);
    const fa=f.found_at?new Date(f.found_at).toLocaleString():'–';
    return '<tr onclick="openFindingByIdx('+idx+')">'+
      '<td style="color:var(--txm);font-family:var(--mono);font-size:9px">'+(i+1)+'</td>'+
      '<td><span class="fc" title="'+esc(f.file)+'">'+esc(sf)+'</span></td>'+
      '<td class="fn">'+f.line+'</td>'+
      '<td>'+snipHTML+'</td>'+
      '<td style="font-family:var(--mono);font-size:9px;color:var(--ac2)">'+esc(f.scriptName||'')+'</td>'+
      '<td style="font-family:var(--mono);font-size:9px;color:var(--txd)">'+esc(f.sourceName||'')+'</td>'+
      '<td>'+confPill(f.scriptConf||0)+'</td>'+
      '<td style="font-family:var(--mono);font-size:9px;color:var(--txm)">'+fa+'</td>'+
    '</tr>';
  }).join('');
}

function openFindingByIdx(i){ if(allFindings[i]) openFinding(allFindings[i]); }

function openFinding(f){
  const sc=scripts.find(s=>s.id===f.scriptId);
  document.getElementById('fm-title').textContent=f.file.split('/').pop()+':'+f.line;
  document.getElementById('fm-file').textContent=f.file;
  document.getElementById('fm-line').textContent=f.line;
  // Script link + desc + hierarchy
  document.getElementById('fm-script-lnk').textContent=f.scriptName||f.scriptId||'–';
  document.getElementById('fm-script-desc').textContent=sc?.description||f.scriptDesc||'';
  const hier=scriptGroupPath(f.scriptId||'');
  document.getElementById('fm-script-hier').textContent=hier.length?'↳ '+hier.join(' | '):'';
  // Source link
  document.getElementById('fm-source-lnk').textContent=f.sourceName||'–';
  // Job link
  const job=jobs.find(j=>j.id===f.jobId);
  document.getElementById('fm-job-lnk').textContent=job?(job.script_name+' → '+job.source_name+' ('+ago(job.created_at)+')'):(f.jobId||'–');
  document.getElementById('fm-found-at').textContent=f.found_at?new Date(f.found_at).toLocaleString():'–';
  const snipEl=document.getElementById('fm-snippet');
  snipEl.innerHTML='';
  // Use snippet_html (ANSI-converted) if available, otherwise fall back to
  // plain snippet text — both go through appendOutputLines so newlines become
  // divs and the layout is always correct.
  const snipSource=f.snippet_html||f.snippet||'(no snippet)';
  appendOutputLines(snipEl, snipSource);
  _fmJobId=f.jobId; _fmScriptId=f.scriptId; _fmSourceId=f.sourceId;
  openModal('findingModal');
}

function fmNavJob(e){e.preventDefault();if(_fmJobId){closeModal('findingModal');openJobDetail(_fmJobId);}return false;}
function fmNavScript(e){e.preventDefault();if(_fmScriptId){closeModal('findingModal');openScriptModal(_fmScriptId);}return false;}
function fmNavSource(e){e.preventDefault();if(_fmSourceId){closeModal('findingModal');openSourceModal(_fmSourceId);}return false;}

// ════════════════════════════════════════════════
// Force Graph
// ════════════════════════════════════════════════
// Graph node key insight: files are shared across jobs.
// A "file node" key is just "file:<path>" — if two jobs have findings
// in the same file, they share the same file node (but have separate edges
// from their respective job nodes).
// Finding nodes remain unique per job+index: "find:<jobIdx>:<findIdx>"
// Dir nodes: "dir:<path>" — also shared across jobs.
//
// Confidence-based finding colors: same palette as conf pills.

const JOB_COLORS=['#28e070','#20c8e0','#e0a020','#b060f8','#e06020','#18d0a0','#7080f0','#e05090'];
const DIR_COLOR='#2a3a50';
const CONF_HI='#28e070'; const CONF_MID='#e0a020'; const CONF_LO='#e84040';

function confColor(v){ return v>=75?CONF_HI:v>=40?CONF_MID:CONF_LO; }

let graphState=null;
function resetGraphView(){ if(graphState) graphState.resetView(); }

function renderGraph(findings){
  const wrap=document.getElementById('graphWrap');
  const canvas=document.getElementById('graphCanvas');
  const dpr=window.devicePixelRatio||1;
  canvas.width=Math.round(wrap.clientWidth*dpr);
  canvas.height=Math.round(wrap.clientHeight*dpr);
  if(graphState){graphState.destroy();graphState=null;}
  if(!findings.length){
    const ctx=canvas.getContext('2d');
    ctx.scale(dpr,dpr);
    ctx.fillStyle='#030405';ctx.fillRect(0,0,wrap.clientWidth,wrap.clientHeight);
    ctx.fillStyle='#3a4e62';ctx.font='10px JetBrains Mono,monospace';
    ctx.textAlign='center';ctx.fillText('No findings',wrap.clientWidth/2,wrap.clientHeight/2);
    return;
  }
  graphState=new ForceGraph(canvas,findings,dpr);
  graphState.start();
  buildGraphLegend(findings);
}

function buildGraphLegend(findings){
  const jobIds=[...new Set(findings.map(f=>f.jobId))];
  let h=jobIds.map((jid,i)=>{
    const col=JOB_COLORS[i%JOB_COLORS.length];
    const name=(findings.find(f=>f.jobId===jid)||{}).scriptName||jid;
    return '<div class="gl-item"><div class="gl-dot" style="background:'+col+'"></div>'+esc(name.length>22?name.slice(0,19)+'…':name)+'</div>';
  }).join('');
  h+='<div class="gl-item"><div class="gl-dot" style="background:'+DIR_COLOR+';border:1px solid #405870"></div>Directory</div>';
  h+='<div class="gl-item"><div class="gl-dot" style="background:'+CONF_HI+'"></div>High conf</div>';
  h+='<div class="gl-item"><div class="gl-dot" style="background:'+CONF_MID+'"></div>Mid conf</div>';
  h+='<div class="gl-item"><div class="gl-dot" style="background:'+CONF_LO+'"></div>Low conf</div>';
  document.getElementById('graphLegend').innerHTML=h;
}

function darken(hex,f){
  const r=parseInt(hex.slice(1,3),16),g=parseInt(hex.slice(3,5),16),b=parseInt(hex.slice(5,7),16);
  return '#'+[r,g,b].map(v=>Math.max(0,Math.round(v*f)).toString(16).padStart(2,'0')).join('');
}

class ForceGraph{
  constructor(canvas,findings,dpr){
    this.canvas=canvas; this.dpr=dpr; this.ctx=canvas.getContext('2d');
    this.findings=findings; this.W=canvas.width/dpr; this.H=canvas.height/dpr;
    this.nodes=[]; this.edges=[]; this.animId=null;
    this.dragging=null; this.hovering=null; this.wasDragging=false;
    this.viewX=0; this.viewY=0; this.viewScale=1;
    this.isPanning=false; this.panStart={x:0,y:0};
    this._build(); this._bindEvents();
  }

  _build(){
    const cx=this.W/2, cy=this.H/2;
    const jobIds=[...new Set(this.findings.map(f=>f.jobId))];
    const nodeMap={}; let ni=0;
    const addNode=(id,props)=>{
      if(nodeMap[id]) return nodeMap[id];
      const a=Math.random()*Math.PI*2, d=30+Math.random()*100;
      const n={id,...props,x:cx+Math.cos(a)*d,y:cy+Math.sin(a)*d,vx:0,vy:0,idx:ni++,fixed:false};
      this.nodes.push(n); nodeMap[id]=n; return n;
    };
    const addEdge=(a,b,len,stiff)=>{
      // avoid duplicate edges
      if(!this.edges.find(e=>(e.a===a&&e.b===b)||(e.a===b&&e.b===a)))
        this.edges.push({a,b,len,stiff});
    };

    for(let ji=0;ji<jobIds.length;ji++){
      const jid=jobIds[ji];
      const jCol=JOB_COLORS[ji%JOB_COLORS.length];
      const jFindings=this.findings.filter(f=>f.jobId===jid);
      const jName=(jFindings[0]||{}).scriptName||jid;
      const jNode=addNode('job:'+jid,{label:jName,type:'job',color:jCol,r:9});
      jNode.x=cx+(ji-jobIds.length/2+.5)*150; jNode.y=cy;

      for(let fi=0;fi<jFindings.length;fi++){
        const f=jFindings[fi];
        const dir=f.file.substring(0,f.file.lastIndexOf('/'))||'/';
        const fname=f.file.split('/').pop();
        const fileCol=darken(jCol,.5);

        // Shared dir node (by path only, across all jobs)
        const dirNode=addNode('dir:'+dir,{label:dir.split('/').pop()||dir,type:'dir',color:DIR_COLOR,r:6});
        addEdge(jNode,dirNode,100,.025);

        // Shared file node (by path only)
        const fileNode=addNode('file:'+f.file,{label:fname,type:'file',color:fileCol,r:5});
        addEdge(dirNode,fileNode,60,.045);
        // If this file is also referenced from a different job, add that edge too
        addEdge(jNode,fileNode,90,.015);

        // Unique finding node
        const fNode=addNode('find:'+jid+':'+fi,{label:'L'+f.line,type:'finding',color:confColor(f.scriptConf||0),r:4,_finding:f});
        addEdge(fileNode,fNode,28,.1);
      }
    }
  }

  _bindEvents(){
    const c=this.canvas;
    this._onMouseDown=e=>{
      const p=this._toWorld(e.offsetX,e.offsetY);
      const hit=this._hitNode(p.x,p.y);
      if(hit){this.dragging=hit;hit.fixed=true;this.wasDragging=false;}
      else{this.isPanning=true;this.panStart={x:e.clientX-this.viewX,y:e.clientY-this.viewY};}
    };
    this._onMouseMove=e=>{
      const p=this._toWorld(e.offsetX,e.offsetY);
      if(this.dragging){this.dragging.x=p.x;this.dragging.y=p.y;this.dragging.vx=0;this.dragging.vy=0;this.wasDragging=true;}
      else if(this.isPanning){this.viewX=e.clientX-this.panStart.x;this.viewY=e.clientY-this.panStart.y;}
      const h=this._hitNode(p.x,p.y);
      this.hovering=h;this._tip(h,e.offsetX,e.offsetY);
    };
    this._onMouseUp=()=>{if(this.dragging){this.dragging.fixed=false;this.dragging=null;}this.isPanning=false;};
    this._onWheel=e=>{
      e.preventDefault();
      const f=e.deltaY<0?1.1:.91;
      this.viewX=(this.viewX-e.offsetX)*f+e.offsetX;
      this.viewY=(this.viewY-e.offsetY)*f+e.offsetY;
      this.viewScale*=f;
    };
    this._onClick=e=>{
      if(this.wasDragging){this.wasDragging=false;return;}
      const p=this._toWorld(e.offsetX,e.offsetY);
      const h=this._hitNode(p.x,p.y);
      if(h?._finding) openFinding(h._finding);
    };
    c.addEventListener('mousedown',this._onMouseDown);
    c.addEventListener('mousemove',this._onMouseMove);
    c.addEventListener('mouseup',this._onMouseUp);
    c.addEventListener('wheel',this._onWheel,{passive:false});
    c.addEventListener('click',this._onClick);
  }

  _toWorld(sx,sy){return{x:(sx-this.viewX)/this.viewScale,y:(sy-this.viewY)/this.viewScale};}
  _hitNode(wx,wy){
    for(let i=this.nodes.length-1;i>=0;i--){
      const n=this.nodes[i]; const dx=n.x-wx,dy=n.y-wy;
      if(dx*dx+dy*dy<=(n.r+5)*(n.r+5)) return n;
    }
    return null;
  }
  _tip(n,sx,sy){
    const t=document.getElementById('graphTip');
    if(!n){t.style.display='none';return;}
    let h='';
    if(n.type==='job') h='<b style="color:'+n.color+'">Job</b> · '+esc(n.label);
    else if(n.type==='dir') h='<b style="color:#6090b0">Dir</b><br>'+esc(n.label);
    else if(n.type==='file') h='<b style="color:'+n.color+'">File</b><br>'+esc(n.label);
    else if(n.type==='finding'){
      const f=n._finding;
      const sc=scripts.find(s=>s.id===f.scriptId);
      const hier=scriptGroupPath(f.scriptId||'');
      h='<b style="color:'+n.color+'">Finding</b> Line '+f.line+'<br>'+esc(f.file.split('/').pop());
      if(sc) h+='<br><span style="color:#7090a0">'+esc(sc.name)+'</span>';
      if(sc?.description) h+='<br><span style="color:#506070;font-size:9px">'+esc(sc.description.slice(0,80))+'</span>';
      if(hier.length) h+='<br><span style="color:#405060;font-size:9px">'+esc(hier[0])+'</span>';
      h+='<br><span style="color:#405060">click to detail</span>';
    }
    t.innerHTML=h; t.style.display='block';
    t.style.left=(sx+14)+'px'; t.style.top=(sy+8)+'px';
  }

  resetView(){this.viewX=0;this.viewY=0;this.viewScale=1;}
  start(){this._tick();}
  _tick(){this.animId=requestAnimationFrame(()=>this._tick());this._sim();this._draw();}

  _sim(){
    const ns=this.nodes,es=this.edges,N=ns.length;
    const rep=1000,damp=.84,dt=.46,cx=this.W/2,cy=this.H/2;
    for(let i=0;i<N;i++){
      const a=ns[i]; if(a.fixed) continue;
      for(let j=i+1;j<N;j++){
        const b=ns[j];
        let dx=a.x-b.x,dy=a.y-b.y,d2=dx*dx+dy*dy;
        if(d2<.5)d2=.5;
        const inv=1/Math.sqrt(d2),f=rep*inv*inv;
        const fx=dx*inv*f,fy=dy*inv*f;
        a.vx+=fx;a.vy+=fy;b.vx-=fx;b.vy-=fy;
      }
      a.vx+=(cx-a.x)*.0016; a.vy+=(cy-a.y)*.0016;
    }
    for(const e of es){
      const a=e.a,b=e.b,dx=b.x-a.x,dy=b.y-a.y;
      const d=Math.sqrt(dx*dx+dy*dy)||1,f=(d-e.len)*e.stiff;
      const fx=dx/d*f,fy=dy/d*f;
      if(!a.fixed){a.vx+=fx;a.vy+=fy;} if(!b.fixed){b.vx-=fx;b.vy-=fy;}
    }
    for(const n of ns){if(n.fixed)continue;n.vx*=damp;n.vy*=damp;n.x+=n.vx*dt;n.y+=n.vy*dt;}
  }

  _draw(){
    const ctx=this.ctx,dpr=this.dpr,W=this.canvas.width/dpr,H=this.canvas.height/dpr;
    ctx.save();ctx.scale(dpr,dpr);
    ctx.fillStyle='#030405';ctx.fillRect(0,0,W,H);
    ctx.translate(this.viewX,this.viewY);ctx.scale(this.viewScale,this.viewScale);

    // Edges — much more visible now
    for(const e of this.edges){
      const isJobEdge=e.a.type==='job'||e.b.type==='job';
      const isFileEdge=e.a.type==='file'||e.b.type==='file';
      ctx.strokeStyle=isJobEdge?'rgba(64,80,110,.75)':isFileEdge?'rgba(50,70,95,.6)':'rgba(42,58,80,.5)';
      ctx.lineWidth=isJobEdge?.9:.65;
      ctx.beginPath();ctx.moveTo(e.a.x,e.a.y);ctx.lineTo(e.b.x,e.b.y);ctx.stroke();
    }

    // Nodes: dir < file < job < finding (top)
    const order=['dir','file','job','finding'];
    for(const type of order){
      for(const n of this.nodes){
        if(n.type!==type) continue;
        const hov=n===this.hovering;
        ctx.globalAlpha=type==='dir'?.55:type==='file'?.7:1;
        ctx.beginPath();ctx.arc(n.x,n.y,n.r+(hov?2:0),0,Math.PI*2);
        ctx.fillStyle=n.color;ctx.fill();
        ctx.globalAlpha=1;
        if(hov||(type==='job')||(type==='finding')){
          ctx.strokeStyle=hov?'#ffffff':n.color;
          ctx.lineWidth=hov?2:type==='job'?1:.8;
          ctx.stroke();
        }
      }
    }

    // Labels
    if(this.viewScale>.45){
      ctx.textAlign='center';
      for(const n of this.nodes){
        if(n.type==='job'){
          ctx.font='bold 10px JetBrains Mono,monospace';
          ctx.fillStyle='rgba(220,230,240,.9)';
          ctx.fillText(n.label.length>16?n.label.slice(0,13)+'…':n.label,n.x,n.y+n.r+11);
        } else if(n.type==='file'&&this.viewScale>.8){
          ctx.font='9px JetBrains Mono,monospace';
          ctx.fillStyle='rgba(140,160,180,.65)';
          ctx.fillText(n.label,n.x,n.y+n.r+9);
        } else if(n.type==='finding'&&this.viewScale>1.5){
          ctx.font='8px JetBrains Mono,monospace';
          ctx.fillStyle='rgba(220,230,240,.5)';
          ctx.fillText('L'+n._finding.line,n.x,n.y+n.r+8);
        }
      }
    }
    ctx.restore();
  }

  destroy(){
    cancelAnimationFrame(this.animId);
    const c=this.canvas;
    ['mousedown','mousemove','mouseup','wheel','click'].forEach(ev=>{
      const k='_on'+ev.charAt(0).toUpperCase()+ev.slice(1);
      c.removeEventListener(ev,this[k],ev==='wheel'?{passive:false}:undefined);
    });
  }
}

// ════════════════════════════════════════════════
// Modals
// ════════════════════════════════════════════════
function openModal(id){document.getElementById(id).classList.add('open');}
function closeModal(id){document.getElementById(id).classList.remove('open');}
document.querySelectorAll('.overlay').forEach(o=>o.addEventListener('click',e=>{if(e.target===o)o.classList.remove('open');}));

// ════════════════════════════════════════════════
// Init
// ════════════════════════════════════════════════
// ════════════════════════════════════════════════
// Reload
// ════════════════════════════════════════════════
async function reloadData(){
  const btn=document.getElementById('reloadBtn');
  btn.disabled=true; btn.textContent='↺ …';
  await api('POST','/api/reload');
  await refreshAll();
  // re-render whatever page is active
  if(currentPage==='dashboard') renderDashJobList();
  else if(currentPage==='jobs') renderJobList();
  else if(currentPage==='scripts') renderScriptTree();
  else if(currentPage==='sources') renderSourceList();
  btn.disabled=false; btn.textContent='↺ Reload';
}

// ════════════════════════════════════════════════
// Timing modal
// ════════════════════════════════════════════════
function openTimingModal(scriptId){
  const sc=scripts.find(s=>s.id===scriptId);
  if(!sc) return;
  document.getElementById('tm-title').textContent='Timing: '+sc.name;
  document.getElementById('tm-script-name').textContent=sc.name;
  document.getElementById('tm-conf').innerHTML=confPill(sc.confidence||0);
  document.getElementById('tm-desc').textContent=sc.description||'–';

  const hasData=sc.timing_samples>0 && sc.avg_bytes_per_sec>0;
  document.getElementById('tm-no-data').style.display=hasData?'none':'';
  document.getElementById('tm-data').style.display=hasData?'grid':'none';

  if(hasData){
    document.getElementById('tm-bps').textContent=fmtBytes(sc.avg_bytes_per_sec)+'/s';
    document.getElementById('tm-lps').textContent=fmtLines(sc.avg_lines_per_sec)+'/s';
    document.getElementById('tm-samples').textContent=sc.timing_samples+' run'+(sc.timing_samples!==1?'s':'');

    // Per-source estimates
    const tbody=document.getElementById('tm-est-body');
    tbody.innerHTML=sources.map(src=>{
      const byBytes=src.size_bytes&&sc.avg_bytes_per_sec>0 ? fmtEst(src.size_bytes/sc.avg_bytes_per_sec) : '–';
      const byLines=src.line_count&&sc.avg_lines_per_sec>0 ? fmtEst(src.line_count/sc.avg_lines_per_sec) : '–';
      const sz=src.size_bytes?fmtBytes(src.size_bytes):'?';
      const lc=src.line_count?fmtLines(src.line_count):'?';
      return '<tr>'+
        '<td style="font-weight:500">'+esc(src.name)+'</td>'+
        '<td style="font-family:var(--mono);font-size:10px;color:var(--txd)">'+sz+'</td>'+
        '<td style="font-family:var(--mono);font-size:10px;color:var(--txd)">'+lc+'</td>'+
        '<td style="font-family:var(--mono);font-size:11px;color:var(--ac3)">'+byBytes+'</td>'+
        '<td style="font-family:var(--mono);font-size:11px;color:var(--ac2)">'+byLines+'</td>'+
      '</tr>';
    }).join('')||'<tr><td colspan="5" style="color:var(--txm);font-family:var(--mono);font-size:10px;padding:8px">No sources defined</td></tr>';
  }

  // Past jobs
  const pastJobs=jobs.filter(j=>j.script_id===scriptId&&j.duration_sec>0);
  const jbody=document.getElementById('tm-jobs-body');
  jbody.innerHTML=pastJobs.length?pastJobs.slice().reverse().slice(0,20).map(j=>{
    const dur=j.duration_sec?j.duration_sec.toFixed(1)+'s':'–';
    const fc=j.findings?.length||0;
    return '<tr onclick="openJobDetail(\''+j.id+'\')" style="cursor:pointer">'+
      '<td style="font-weight:500">'+esc(j.source_name||j.source_id)+'</td>'+
      '<td>'+badge(j.status)+'</td>'+
      '<td style="font-family:var(--mono);font-size:11px;color:var(--ac3)">'+dur+'</td>'+
      '<td style="font-family:var(--mono);font-size:11px;color:'+(fc?'var(--red)':'var(--txm)')+'">'+fc+'</td>'+
      '<td style="font-family:var(--mono);font-size:9px;color:var(--txm)">'+ago(j.created_at)+'</td>'+
    '</tr>';
  }).join(''):'<tr><td colspan="5" style="color:var(--txm);font-family:var(--mono);font-size:10px;padding:8px">No completed runs yet</td></tr>';

  openModal('timingModal');
}

async function init(){
  await refreshAll();
  renderDashJobList();
  fetchMetrics();
  setInterval(async()=>{
    if(currentPage==='dashboard'||currentPage==='jobs'){
      jobs=await api('GET','/api/jobs')||[];
      buildAllFindings(); updateStats();
      if(currentPage==='dashboard') renderDashJobList();
      else renderJobList();
    }
    if(currentPage==='dashboard') fetchMetrics();
  },4000);
  window.addEventListener('resize',()=>{
    if(currentPage==='findings'&&activeVTab==='graph') renderGraph(getFilteredFindings());
  });
}
init();
</script>
</body>
</html>`

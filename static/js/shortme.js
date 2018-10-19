function DisplayQR(forWho, content) {
    jQuery("#" + forWho).qrcode({
        text: content,
        width: 128,
        height: 128
    });
}

function Short(longURLID) {
    var longURL = document.getElementById(longURLID).value;
    longURL = longURL.trim();
    if (longURL === "") {
        alert("Input text is empty. :-)");
        return
    }

    var shortURL = "";

    jQuery.ajax ({
        url: "/short",
        type: "POST",
        data: JSON.stringify({longURL: longURL}),
        dataType: "json",
        contentType: "application/json",
        success: function(data, textStatus, xhr){
            shortURL = data.shortURL;
        },
        error: function(jqXHR, textStatus, errorThrown) {
            alert(JSON.parse(jqXHR.responseText).msg);
        }
    }).always(function() {
        if (shortURL === "") {
            return
        }

        // add blank lines
        var blankList = document.getElementById("shortURLBlankLine");
        if (!blankList.hasChildNodes()) {
            var lineBreak = document.createElement("br");
            document.getElementById("shortURLBlankLine").appendChild(lineBreak);
        }

        // specify the shortened url
        document.getElementById("shortenedURL").innerHTML = shortURL;

        // add shortened qr code
        document.getElementById("shortenedQR").innerHTML = "";
        DisplayQR("shortenedQR", shortURL);

        // add shortened url preview
        var shortenedURLPreviewIframe = document.getElementById("shortenedURLPreviewIframe");
        if (shortenedURLPreviewIframe === null) {
            var preview = document.createElement("iframe");
            preview.setAttribute("src", longURL);
            preview.setAttribute("frameBorder", 0);
            preview.setAttribute("scrolling", "auto");
            preview.setAttribute("id", "shortenedURLPreviewIframe");
            preview.setAttribute("sandbox", "");
            preview.setAttribute("security", "restricted");
            document.getElementById("shortenedURLPreview").appendChild(preview);
        } else {
            shortenedURLPreviewIframe.setAttribute("src", longURL);
        }
    })
}

function Expand(shortURLID) {
    var shortURL = document.getElementById(shortURLID).value;
    shortURL = shortURL.trim();
    if (shortURL === "") {
        alert("Input text is empty. :-)");
        return
    }

    var longURL = "";

    jQuery.ajax ({
        url: "/expand",
        type: "POST",
        data: JSON.stringify({shortURL: shortURL}),
        dataType: "json",
        contentType: "application/json",
        success: function(data, textStatus, xhr){
            longURL = data.longURL;
        },
        error: function(jqXHR, textStatus, errorThrown) {
            alert(JSON.parse(jqXHR.responseText).msg);
        }
    }).always(function() {
        if (longURL === "") {
            return
        }

        // clear blank lines
        var blankList = document.getElementById("expandedURLBlankLine");
        if (!blankList.hasChildNodes()) {
            var lineBreak = document.createElement("br");
            document.getElementById("expandedURLBlankLine").appendChild(lineBreak);
        }

        // specify the expanded url
        document.getElementById("expandedURL").innerHTML = longURL;


        // clear expanded qr code
        document.getElementById("expandedQR").innerHTML = "";
        DisplayQR("expandedQR", longURL);

        // add expanded url preview
        var expandedURLPreviewIframe = document.getElementById("expandedURLPreviewIframe");
        if (expandedURLPreviewIframe === null) {
            var preview = document.createElement("iframe");
            preview.setAttribute("src", longURL);
            preview.setAttribute("frameBorder", 0);
            preview.setAttribute("scrolling", "auto");
            preview.setAttribute("id", "expandedURLPreviewIframe");
            preview.setAttribute("sandbox", "");
            preview.setAttribute("security", "restricted");
            document.getElementById("expandedURLPreview").appendChild(preview);
        } else {
            expandedURLPreviewIframe.setAttribute("src", longURL);
        }
    })
}

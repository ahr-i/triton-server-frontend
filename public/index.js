$(function() {
    let tabBtn1 = $('#tab-btn-stable');
    let dropdown1 = $('#dropdown-menu-model');
    let submit = $('#submit')
    let textInput = $('#text-input')
    let imageDisplay = $('#image-display')
    let loading = $('#loading')

    $.get('/get/model-list', function(response) {
        $.each(response, function(_, model) {
            let row = "<option value='" + model + "'>" + model + "</option>";

            dropdown1.append(row);
        });
    });

    $(document).ready(function() {
        tabBtn1.click(function() {
            $.get('/get/url/stable_diffusion', function(response) {
                url = response.url;

                window.location.href = url;
            });
        });
    });

    $(document).ready(function() {
        submit.click(function() {
            let prompt = textInput.val();
            let model = dropdown1.val();
            let data = {
                prompt: prompt,
                seed: "1"
            };

            if (model == "Select a model...") {
                alert("Please select a model");

                return;
            } else if (!prompt.trim()) {
                alert("Please enter a prompt.");

                return;
            }

            imageDisplay.empty();
            loading.text("Loading...")

            fetch("/model/" + model + "/infer", {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            })
            .then(response => response.blob())
            .then(blob => {
                let imageUrl = URL.createObjectURL(blob);
                let imgTag = $("<img>", { src: imageUrl, alt: "Generated Image" });
                imageDisplay.empty().append(imgTag);

                loading.text("Finish")
            })
            .catch(error => {
                console.error("ERROR" + error);

                loading.text("Error")
            });
        });
    });
});